package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	Reports map[string]*Report
)

type Report struct {
	ReportID   string
	Score      int64
	PlayerName string
}

func init() {
	Reports = make(map[string]*Report)
	Reports["hjkhsbnmn123"] = &Report{"hjkhsbnmn123", 100, "astaxie"}
	Reports["mjjkxsxsaa23"] = &Report{"mjjkxsxsaa23", 101, "someone"}
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/finance_report?charset=utf8", 30)

}

func AddOne(Report Report) (ReportID string) {
	Report.ReportID = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Reports[Report.ReportID] = &Report
	return Report.ReportID
}

func GetOne(ReportID string) (Report *Report, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT name FROM user  ").Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps) // slene
	}

	if v, ok := Reports[ReportID]; ok {
		return v, nil
	}
	return nil, errors.New("ReportID Not Exist")
}

func GetAll() map[string]*Report {
	return Reports
}

func Update(ReportID string, Score int64) (err error) {
	if v, ok := Reports[ReportID]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ReportID Not Exist")
}

func Delete(ReportID string) {
	delete(Reports, ReportID)
}
