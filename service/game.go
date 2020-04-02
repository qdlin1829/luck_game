package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"luck_game/config"
	"luck_game/model"
	"strings"
	"time"
)

type GameService struct {
}

type gameData struct {
	CurrentOpenDateTime time.Duration `json:"currentOpenDateTime"`
	OpenDateTime        time.Duration `json:"openDateTime"`
	OpenedCount         int           `json:"openedCount"`
	DailyTotal          int           `json:"dailyTotal"`
	GameCode            string        `json:"gameCode"`
	OpenNum             []int         `json:"openNum"`
	PreIssue            string        `json:"preIssue"`
	Issue               string        `json:"issue"`
	SumArr              []int         `json:"sumArr"`
}

type gameOpen struct {
	Tc    int         `json:"tc"`    // 总开奖数
	Dc    int         `json:"dc"`    // 开奖数
	Di    string      `json:"di"`    // 下一期
	Pdi   string      `json:"pdi"`   // 下一期
	Pdc   interface{} `json:"pdc"`   // 开奖数
	Sbs   string      `json:"sbs"`   // 总合。大小
	Ssd   string      `json:"ssd"`   // 总合。单双
	SumFS int         `json:"sumfs"` // 总合
	ln    string      `json:"ln"`    //  游戏名称
	Dt    int64       `json:"dt"`    // 开奖日期
}

func (g *GameService) Set() (ret interface{}, err error) {

	data := gameData{}
	url := "https://www.131313.com/data/Current/xyft/CurIssue.json?" + strings.ToUpper(config.Md5(string(time.Now().Unix())))

	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	// 返回数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	json.Unmarshal([]byte(string(body)), &data)
	if data.Issue == "" {
		return ret, err
	}

	var str = []string{}
	for i := 0; i < len(data.OpenNum); i++ {
		str = append(str, fmt.Sprintf("%d", data.OpenNum[i]))
	}

	xyft_key := "xyft"
	xyftRet, _ := rdb.Get(xyft_key).Result()
	xyft_res := &model.Xyft{}
	json.Unmarshal([]byte(xyftRet), xyft_res)

	if xyft_res.Issue == data.PreIssue {
		return xyft_res, nil
	}

	// 存入reids
	var xyftOpen = make(map[string]interface{})
	xyftOpen["issue"] = data.PreIssue
	xyftOpen["sumfs"] = data.SumArr[0]
	xyftOpen["ds"] = dsSwtich(data.SumArr[1])
	xyftOpen["number"] = data.OpenNum
	xyftOpen["dx"] = dxSwtich(data.SumArr[2])
	xyftOpen["new_issue"] = data.Issue
	xyftOpen["new_open_time"] = int64(data.OpenDateTime / 1000)

	open, _ := json.Marshal(xyftOpen)
	rdb.Set(xyft_key, open, 500*time.Second).Result()

	// 写入db
	xyft := model.Xyft{
		Issue:       data.PreIssue,
		Sumfs:       data.SumArr[0],
		Ds:          dsSwtich(data.SumArr[1]),
		Dx:          dxSwtich(data.SumArr[2]),
		Number:      strings.Join(str, ","),
		NewIssue:    data.Issue,
		NewOpenTime: int64(data.OpenDateTime / 1000),
		CreateTime:  time.Now().Unix(),
	}

	Db.Table("go_xyft").Insert(xyft)

	return xyftOpen, nil
}

func dsSwtich(i int) (str string) {

	switch i {
	case 1:
		str = "单"
	case 2:
		str = "双"
	default:
	}
	return str
}

func dxSwtich(i int) (str string) {

	switch i {
	case 1:
		str = "大"
	case 2:
		str = "小"
	default:
	}
	return str
}
