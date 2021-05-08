package core

import (
	"BandBinance/config"
	"context"
	"strconv"
	"strings"
)

func Hours24Tickers() float64 {
	res, err := client.NewListPriceChangeStatsService().Symbol(config.Symbol).Do(context.Background())
	if err != nil {
		return 0
	}
	c, _ := strconv.ParseFloat(res[0].PriceChangePercent, 64)
	return c
}

func InitSaveData() {
	res, err := client.NewListPriceChangeStatsService().Symbol(config.Symbol).Do(context.Background())
	if err != nil {
		return
	}
	weightedAvgPrice, _ := strconv.ParseFloat(res[0].WeightedAvgPrice, 64)
	bs := []Bet{}
	rightSize := len(strings.Split(res[0].WeightedAvgPrice, ".")[1])
	for k, v := range config.NetRa {
		bp := round(weightedAvgPrice*(1-v/100), rightSize)
		sp := round(weightedAvgPrice*(1+v/100), rightSize)
		te := Bet{
			BuyPrice:  bp,
			SellPrice: sp,
			Step:      0,
			Type:      k,
		}
		bs = append(bs, te)
	}
	p := PriceData{
		Bs:bs,
		SiL:SimulateBalance{
			Money:100,
			Coin:20,
		},
		Spend:5,
		SetupPrice:weightedAvgPrice,
		O:Ori{
			OriMoney:100,
			OriCoin:20,
		},
		LimitQ:1,
	}
	p.save()

}