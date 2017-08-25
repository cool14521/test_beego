package models

import "github.com/astaxie/beego/logs"

func LogInfo() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"info.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":365}`)

}
