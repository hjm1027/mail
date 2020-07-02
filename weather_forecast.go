package main

import (
	"encoding/json"
	"net/smtp"
	"net/http"
	"io/ioutil"
	"strings"
	"time"
	"fmt"
)

const (
	url  = "https://restapi.amap.com/v3/weather/weatherInfo?"
	key  = "d1c5947ba4cb0facaafae2cc5a4d532c"
	city = "420111"
	exten  = "all" 
	output = "JSON"
)

type Out struct{
	Status string `json:"status"`
	Count string  `json:""count`
	Info  string  `json:"info"`
	Infocode string `json:"infocode"`
	Forecasts  []Middle `json:"forecasts"`
}

type Middle struct{
	City string `json:"city"`
	Adcode string `json:"adcode"`
	Province string `json:"province"`
	Reporttime string `json:"reporttime"`
	Casts  []In  `json:"casts"`
}

type In struct{
	Date string `json:"date"`
	Week string `json:"week"`
	Dayweather string `json:"dayweather"`
	Nightweather string `json:"nightweather"`
	Daytemp string `json:"daytemp"`
	Nighttemp string `json:"nighttemp"`
	Daywind string `json:"daywind"`
	Nightwind string `json:"nightwind"`
	Daypower string `json:"daypower"`
	Nightpower string `json:"nightpower"`
}

// 邮件信息
func sendEmail(subject, body string) {
	user := "1749511631@qq.com"
	pwd := "****************"
	host := "smtp.qq.com:25"
	to := "1749511631@qq.com" //可以用;隔开发送多个
	fmt.Println("send email")
	err := sendToMail(user, pwd, host, to, subject, body)
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

// 发送邮件
func sendToMail(user, password, host, to, subject, body string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])

	msg := []byte("To:" + to + "\nFrom: " + user + "<" +user + ">\nSubject: " + subject + "\n" + "\n\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}


func main(){
	var a Out
	longurl :=url + "key=" + key + "&city=" + city + "&extensions=" + exten + "&output" + output 

	resp, err := http.Get(longurl)
	if err != nil {
		fmt.Printf("网络请求失败 %v", err)
		return
	}
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	
	err = json.Unmarshal([]byte(string(rbody)), &a)
	if err!= nil {
		fmt.Printf("Json解析失败 %v", err)
		return
	}
	
	//编辑邮件内容
	b := a.Forecasts[0]
	c := b.Casts[0]
	var subject, body string
	subject = b.Province + "省" + b.City + "今日天气(" + c.Date + ")"
	body = "白天天气：" + c.Dayweather + "，温度："  + c.Daytemp + "℃  " + c.Daywind + "风" + c.Daypower + "级。\n" +
	"夜晚天气：" + c.Nightweather + "，温度："  + c.Nighttemp + "℃  " + c.Nightwind + "风" + c.Nightpower + "级。\n"

	//计时器
	d := time.Duration(time.Minute)
	t := time.NewTicker(d)
	defer t.Stop()
	for {
		currentTime := time.Now()
		if currentTime.Hour() == 9 {
			sendEmail(subject, body)
			time.Sleep( 23 * time.Hour)
		}
		<-t.C
	}
}
