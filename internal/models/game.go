package models

type Result struct {
	Key string
	Val int64
}

type Game struct {
	Name string
	Code string
	Func string
	Res  []*Result
}

func GetGames() []Game {
	data := []Game{
		{
			Name: "大小", Code: "dx", Func: "Daxiao", Res: []*Result{
				{Key: "d", Val: 0},
				{Key: "x", Val: 0},
			},
		},
		{
			Name: "单双", Code: "ds", Func: "Danshuang", Res: []*Result{
				{Key: "d", Val: 0},
				{Key: "s", Val: 0},
			},
		},
	}

	return data
}
