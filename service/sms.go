package service

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	smsUrl  = "http://101.227.68.49:7891/mt?"
	account = "10690425"
	password = "gzzj1114"
	rf= 2
	tf     = 3
	dc     = 15
	rd     = 1
)

type SmsReps struct {
	Succes bool `json:"success"`
	Id string` json:"id"`
}

type SmsService struct {

}


func (s *SmsService) SmsSend (mobile, msg string) (ret bool, err error) {

	key := fmt.Sprintf("sms_total_%s", mobile)
	num,err := rdb.Get(key).Int()
	if num > 5 {
		return false, errors.New("超操作太频繁,稍后再试")
	}


	err = rdb.Incr(key).Err()
	if err != nil {
		return false, err
	}
	rdb.Expire(key, 2*60*60*time.Second)


	url := fmt.Sprintf("dc=%d&un=%s&pw=%s&tf=%d&rf=%d&rd=%d&da=%s&sm=%s", dc, account,password, tf, rf, rd, mobile, "【糯米雅】:"+msg)
	// get请求
	resp, err := http.Get(smsUrl+url)
	if err != nil {
		return false, err
	}

	defer  resp.Body.Close()
	// 返回数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false,err
	}

	// josn转struct
	res := &SmsReps{}
	err = json.Unmarshal([]byte(string(body)), res)
	if err != nil {
		return false, err
	}

	return res.Succes, nil
}