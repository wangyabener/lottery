package reminder

import (
	"fmt"
	"lottery/internal/models"
	"lottery/internal/services"
	"lottery/utils"
)

func Long() {
	// get lotteries from database
	lotteries := models.GetLotteries()

	// get games from database or config
	games := models.GetGames()

	// generate secends
	secs := [1]int{}
	for i, j := 0, 0; i < len(secs); i++ {
		secs[i] = j
		j = j + 3
	}

	data := make(map[int]map[string]models.Game)

	for _, lottery := range lotteries {
		for _, sec := range secs {
			items := make(map[string]models.Game)
			for _, game := range games {
				// reflect
				f, _ := utils.InvokeObjectMethod(new(services.Lucky), game.Func, lottery.BlockNum)
				s := f[0].String()

				results := make([]*models.Result, 0, len(game.Res))
				for _, res := range game.Res {
					if s == res.Key {
						res.Val++
					}

					results = append(results, &models.Result{
						Key: res.Key,
						Val: res.Val,
					})
				}
				items[game.Code] = models.Game{
					Code: game.Code,
					Res:  results,
				}
			}
			data[sec] = items
		}
	}

	for sec, items := range data {
		fmt.Printf("第 %d 秒", sec)
		fmt.Println("")
		for _, item := range items {
			for _, v := range item.Res {
				fmt.Println(v.Key, v.Val)
			}
		}
	}
}
