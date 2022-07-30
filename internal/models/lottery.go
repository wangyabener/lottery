package models

type Lottery struct {
	Id        int64
	BlockNum  string
	BlockInfo Block
	DaXiao    int64
	DanShuang int64
	NiuNiu    int64
}

func GetLotteries() [10]Lottery {
	data := [10]Lottery{
		{Id: 1, BlockNum: "132aaa", BlockInfo: Block{BlockTime: 1658976847}, DaXiao: 1},
		{Id: 2, BlockNum: "135bbb", BlockInfo: Block{BlockTime: 1658976847}, DaXiao: 2},
		{Id: 3, BlockNum: "133ccc", BlockInfo: Block{BlockTime: 1658976847}, DaXiao: 2},
		{Id: 4, BlockNum: "138ddd", BlockInfo: Block{BlockTime: 1658976847}, DaXiao: 1},
		{Id: 5, BlockNum: "136eee", BlockInfo: Block{BlockTime: 1658976847}, DaXiao: 1},
		{Id: 6, BlockNum: "131fff", BlockInfo: Block{BlockTime: 1658976847}, DaXiao: 2},
		{Id: 7, BlockNum: "138ggg", BlockInfo: Block{BlockTime: 1658976847}, DaXiao: 1},
		{Id: 8, BlockNum: "137hhh", BlockInfo: Block{BlockTime: 1658976847}, DaXiao: 1},
		{Id: 9, BlockNum: "135kkk", BlockInfo: Block{BlockTime: 1658976847}, DaXiao: 1},
		{Id: 10, BlockNum: "139ll", BlockInfo: Block{BlockTime: 1658976847}, DaXiao: 1},
	}

	return data
}
