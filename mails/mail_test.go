package mails_test

import (
	"testing"
	. "../mails"
)

func TestMail(t *testing.T) {
	config := `{"username":"project@wacai.com","password":"123456","host":"mail.wacai.com","port":25}`
	mail := NewEMail(config)
	if mail.Username != "project@wacai.com" {
		t.Fatal("email parse get username error")
	}
	if mail.Password != "123456" {
		t.Fatal("email parse get password error")
	}
	if mail.Host != "mail.wacai.com" {
		t.Fatal("email parse get host error")
	}
	if mail.Port != 25 {
		t.Fatal("email parse get port error")
	}
	mail.To = []string{"buquan@wacai.com"}
	mail.From = "project@wacai.com"
	mail.Subject = "hi, just from wc_utils!"
	mail.Text = "Text Body is, of course, supported!"
	mail.HTML = "<h1>Fancy Html is supported, too!</h1>"
	mail.AttachFile("/Users/astaxie/github/wc_utils/wc_utils.go")
	mail.Send()
}