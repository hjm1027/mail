package main

import (
	"fmt"
	"strings"
	"net/smtp"
)

// 发送邮件
func sendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])

	var content_type string
	if mailtype == "html" {
		content_type = "Content_Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content_Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To:" + to + "\nFrom: " + user + "<" +
        user + ">\nSubject: " + subject + "\n" +
         content_type + "\n\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}


func sendEmail(subject, body string) {
	user := "1749511631@qq.com"
	pwd := "****************"
	host := "smtp.qq.com:25"
	to := "1749511631@qq.com" //可以用;隔开发送多个
	fmt.Println("send email")
	err := sendToMail(user, pwd, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

func main(){
	sendEmail("zsbd","zsbd")
}
