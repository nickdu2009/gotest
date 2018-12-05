package player

import "github.com/nickdu2009/gotest/copy/deepcopy"

// 所有球员的配置信息
type playerPropertyRecord struct {
	ID            int32  `csv:"ID" index:"true"`
	Name          string `csv:"name"`
	Exchange_res3 int32  `csv:"exchange_res3"`
	PositionType  int32  `csv:"position_type"`
	Quality       int32  `csv:"quality"`
	GoodAt        string `csv:"good_at"`
	CardGrid      int32  `csv:"cardgrid_config"`

	// 以下是计算阵容评分使用字段
	Strong        int32   `csv:"strong"`
	Sp            int32   `csv:"sp"`
	Accelerate    int32   `csv:"accelerate"`
	Speed         int32   `csv:"speed"`
	Jump          int32   `csv:"jump"`
	Flexible      int32   `csv:"flexible"`
	Blance        int32   `csv:"balance"`
	Tackling      int32   `csv:"tackling"`
	Dribble       int32   `csv:"dribble"`
	Controlball   int32   `csv:"controlball"`
	Mark          int32   `csv:"mark"`
	Steals        int32   `csv:"steals"`
	Passcross     int32   `csv:"passcross"`
	Passshort     int32   `csv:"passshort"`
	Passlong      int32   `csv:"passlong"`
	Shoot         int32   `csv:"shoot"`
	Shootlong     int32   `csv:"shootlong"`
	Shoothigh     int32   `csv:"shoothigh"`
	Shootstrong   int32   `csv:"shootstrong"`
	Redian        int32   `csv:"radian"`
	Head          int32   `csv:"head"`
	Freekick      int32   `csv:"freekick"`
	Penalty       int32   `csv:"penalty"`
	Tactical      int32   `csv:"tactical"`
	Station       int32   `csv:"station"`
	View          int32   `csv:"view"`
	Reaction      int32   `csv:"reaction"`
	Porgress      int32   `csv:"porgress"`
	Gkbigfoot     int32   `csv:"gk_bigfoot"`
	Gksaveball    int32   `csv:"gk_saveball"`
	Gkcontrolball int32   `csv:"gk_controlball"`
	Gkstation     int32   `csv:"gk_station"`
	Gkreaction    int32   `csv:"gk_reaction"`
	SameID        []int32 `csv:"same_id"`
}

var player = &playerPropertyRecord{
	ID:            1,
	Name:          "小明",
	Exchange_res3: 1,
	PositionType:  2,
	Quality:       3,
	GoodAt:        "origin",
	CardGrid:      4,
	Strong:        4,
	Sp:            6,
	Accelerate:    7,
	Speed:         8,
	Jump:          9,
	Flexible:      10,
	Blance:        11,
	Tackling:      12,
	Dribble:       13,
	Controlball:   14,
	Mark:          18,
	Steals:        19,
	Passcross:     20,
	Passshort:     21,
	Passlong:      22,
	Shoot:         23,
	Shootlong:     24,
	Shoothigh:     25,
	Shootstrong:   26,
	Redian:        24,
	Head:          28,
	Freekick:      29,
	Penalty:       1,
	Tactical:      3,
	Station:       2,
	View:          4,
	Reaction:      4,
	Porgress:      5,
	Gkbigfoot:     4,
	Gksaveball:    6,
	Gkcontrolball: 7,
	Gkstation:     8,
	Gkreaction:    9,
	SameID:        []int32{1, 2, 3},
}

func GetOriginPlayer() *playerPropertyRecord {
	return player
}

func GetPlayerWithCopy() *playerPropertyRecord {
	var newPlayer = &playerPropertyRecord{}
	deepcopy.Copy(newPlayer, player)
	return newPlayer
}

func GetPlayerWithClone() *playerPropertyRecord {
	newPlayer, _ := deepcopy.Clone(player)
	return newPlayer.(*playerPropertyRecord)
}
