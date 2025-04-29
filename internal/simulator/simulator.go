package simulator

import (
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"piggy-bank/internal/engine"
	"piggy-bank/internal/rng"

	"github.com/schollz/progressbar/v3"
)

type SimulationResult struct {
	Game        string   `xlsx:"Game"`
	Count       int64    `xlsx:"Count"`
	Wager       int64    `xlsx:"Wager"`
	Spent       *big.Int `xlsx:"Spent"`
	MaxExposure int64    `xlsx:"Max Exposure"`

	BaseAwardCount int64 `xlsx:"Base Award Count"`

	X1Count   int64 `xlsx:"X1 Count"`
	X10Count  int64 `xlsx:"X10 Count"`
	X100Count int64 `xlsx:"X100 Count"`

	BaseAward *big.Int `xlsx:"Base Award"`
	Award     *big.Int `xlsx:"Award"`

	BaseAwardSquareSum *big.Int `xlsx:"Base Award Square Sum"`
	AwardSquareSum     *big.Int `xlsx:"Award Square Sum"`

	BaseAwardStandardDeviation *big.Float `xlsx:"Base Award Standard Deviation"`
	AwardStandardDeviation     *big.Float `xlsx:"Award Standard Deviation"`

	Volatility *big.Float `xlsx:"Volatility"`

	RTP         float64 `xlsx:"RTP"`
	RTPBaseGame float64 `xlsx:"RTP Base Game"`
}

type SimulationView struct {
	Game        string `json:"game" xlsx:"Game"`
	Count       string `json:"count" xlsx:"Count"`
	Wager       string `json:"wager" xlsx:"Wager"`
	Spent       string `json:"spent" xlsx:"Spent"`
	MaxExposure string `json:"max_exposure" xlsx:"Max Exposure"`

	AwardCount string `json:"award_count" xlsx:"Award Count"`
	AwardRate  string `json:"award_rate" xlsx:"Award Rate (Hit Rate)"`

	X1Count   string `json:"x1_count" xlsx:"X1 Count"`
	X10Count  string `json:"x10_count" xlsx:"X10 Count"`
	X100Count string `json:"x100_count" xlsx:"X100 Count"`

	X1Rate   string `json:"x1_rate" xlsx:"X1 Rate"`
	X10Rate  string `json:"x10_rate" xlsx:"X10 Rate"`
	X100Rate string `json:"x100_rate" xlsx:"X100 Rate"`

	BaseAward string `json:"base_award" xlsx:"Base Award"`
	Award     string `json:"award" xlsx:"Award"`

	BaseAwardSquareSum string `json:"base_award_square_sum" xlsx:"Base Award Square Sum"`
	AwardSquareSum     string `json:"award_square_sum" xlsx:"Award Square Sum"`

	BaseAwardStandardDeviation float64 `json:"base_award_standard_deviation" xlsx:"Base Award Standard Deviation"`
	AwardStandardDeviation     float64 `json:"award_standard_deviation" xlsx:"Award Standard Deviation"`

	Volatility float64 `json:"volatility" xlsx:"Volatility"`
	RTP        string  `json:"rtp" xlsx:"RTP"`
}

func Simulate(game string, count int64, wager int64, workersCount int, rngService *rng.Service) (*SimulationResult, error) {
	res := &SimulationResult{
		Wager: wager,
		Count: count,
		Game:  game,

		BaseAward: new(big.Int),
		Award:     new(big.Int),
		Spent:     new(big.Int),

		BaseAwardSquareSum: new(big.Int),
		AwardSquareSum:     new(big.Int),

		BaseAwardStandardDeviation: new(big.Float),
		AwardStandardDeviation:     new(big.Float),
	}

	type result struct {
		Wager     int64
		BaseAward int64
	}

	now := time.Now()
	bar := progressbar.NewOptions64(count,
		progressbar.OptionThrottle(200*time.Millisecond),
		progressbar.OptionSetDescription("Simulating..."),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(50),
		progressbar.OptionShowIts(),
		progressbar.OptionOnCompletion(func() {
			fmt.Println("\nTime elapsed:", time.Since(now))
		}),
	)

	rngClient := rngService.GetClient()
	spinFactory := engine.NewSpinFactoryWithAllReelsets(rngClient)

	inputCh := make(chan int64, workersCount)
	outputCh := make(chan result, workersCount)
	errCh := make(chan error, 1)

	wg := new(sync.WaitGroup)

	worker := func(wg *sync.WaitGroup, inputCh <-chan int64, outputCh chan<- result) {
		defer wg.Done()

		for {
			if _, ok := <-inputCh; !ok {
				return
			}

			spin, err := spinFactory.Generate(wager)
			if err != nil {
				errCh <- err
				return
			}

			outputCh <- result{
				Wager:     spin.Wager,
				BaseAward: spin.Award,
			}
		}
	}

	go func() {
		defer close(inputCh)
		for i := int64(0); i < count; i++ {
			inputCh <- i
		}
	}()

	go func() {
		for i := 0; i < workersCount; i++ {
			wg.Add(1)
			go worker(wg, inputCh, outputCh)
		}
		wg.Wait()
		close(outputCh)
	}()

	i := 0

Loop:
	for {
		select {
		case output, ok := <-outputCh:
			if !ok {
				break Loop
			}

			award := output.BaseAward

			res.BaseAward.Add(res.BaseAward, big.NewInt(output.BaseAward))
			res.Award.Add(res.Award, big.NewInt(award))
			res.Spent.Add(res.Spent, big.NewInt(output.Wager))

			if award > res.MaxExposure {
				res.MaxExposure = award
			}

			if award > 0 {
				res.BaseAwardCount++
			}

			if award >= wager*1 {
				res.X1Count++
			}

			if award >= wager*10 {
				res.X10Count++
			}

			if award >= wager*100 {
				res.X100Count++
			}

			res.BaseAwardSquareSum.Add(res.BaseAwardSquareSum, big.NewInt(0).Mul(big.NewInt(output.BaseAward), big.NewInt(output.BaseAward)))
			res.AwardSquareSum.Add(res.AwardSquareSum, big.NewInt(0).Mul(big.NewInt(award), big.NewInt(award)))

			_ = bar.Add(1)
			i++
		case err := <-errCh:
			return nil, err
		}
	}

	baseMeanB := new(big.Float).SetInt(res.BaseAward)
	totalMeanB := new(big.Float).SetInt(res.Award)

	baseMeanB.Quo(baseMeanB, big.NewFloat(float64(res.Count)))
	totalMeanB.Quo(totalMeanB, big.NewFloat(float64(res.Count)))

	res.BaseAwardStandardDeviation = StandardDeviation(res.BaseAwardSquareSum, res.BaseAward, baseMeanB, res.Count)
	res.AwardStandardDeviation = StandardDeviation(res.AwardSquareSum, res.Award, totalMeanB, res.Count)

	res.Volatility = new(big.Float).Quo(res.AwardStandardDeviation, new(big.Float).SetInt64(wager))

	awardF := new(big.Float).SetInt(res.Award)
	baseF := new(big.Float).SetInt(res.BaseAward)
	spentF := new(big.Float).SetInt(res.Spent)

	res.RTP, _ = new(big.Float).Quo(awardF, spentF).Float64()
	res.RTPBaseGame, _ = new(big.Float).Quo(baseF, spentF).Float64()

	return res, nil
}

func (r SimulationResult) View() *SimulationView {
	return &SimulationView{
		Game:        r.Game,
		Count:       fmt.Sprint(r.Count),
		Wager:       fmt.Sprint(r.Wager),
		Spent:       fmt.Sprint(r.Spent),
		MaxExposure: fmt.Sprint(r.MaxExposure),

		AwardCount: fmt.Sprint(r.BaseAwardCount),
		AwardRate:  countToRate(r.BaseAwardCount, r.Count),

		X1Count:   fmt.Sprint(r.X1Count),
		X10Count:  fmt.Sprint(r.X10Count),
		X100Count: fmt.Sprint(r.X100Count),

		X1Rate:   countToRate(r.X1Count, r.Count),
		X10Rate:  countToRate(r.X10Count, r.Count),
		X100Rate: countToRate(r.X100Count, r.Count),

		BaseAward: r.BaseAward.String(),
		Award:     r.Award.String(),

		BaseAwardSquareSum: r.BaseAwardSquareSum.String(),
		AwardSquareSum:     r.AwardSquareSum.String(),

		BaseAwardStandardDeviation: float64FromBigFloat(r.BaseAwardStandardDeviation, 3),
		AwardStandardDeviation:     float64FromBigFloat(r.AwardStandardDeviation, 3),

		Volatility: float64FromBigFloat(r.Volatility, 3),
		RTP:        floatWithPrecision(r.RTP),
	}
}

func StandardDeviation(squareSum, award *big.Int, mean *big.Float, count int64) *big.Float {
	f1 := new(big.Float).SetInt(squareSum)

	f2 := big.NewFloat(-2)
	f2.Mul(f2, mean)
	f2.Mul(f2, new(big.Float).SetInt(award))

	f3 := new(big.Float).SetInt64(count)
	f3.Mul(f3, mean)
	f3.Mul(f3, mean)

	sd := big.NewFloat(0)
	sd.Add(sd, f1)
	sd.Add(sd, f2)
	sd.Add(sd, f3)

	sd.Quo(sd, new(big.Float).SetInt64(count))
	sd.Sqrt(sd)

	return sd
}

func float64FromBigFloat(float *big.Float, i int) float64 {
	value, _ := float.Float64()
	round := math.Pow(10, float64(i))
	return math.Round(value*round) / round
}

func countToRate(count, total int64) string {
	div := float64(count) / float64(total)
	return floatWithPrecision(div) + "%"
}

func floatWithPrecision(f float64) string {
	return fmt.Sprintf("%.3f", f*100)
}
