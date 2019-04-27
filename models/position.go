package models

import (
	"bishe/backend/conf"
	"fmt"

	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

const (
	CITY_SQL = `SELECT city, COUNT(city) AS num FROM position_info WHERE position LIKE ? GROUP BY city`
	SALARY_SQL = `SELECT COUNT(*) AS num FROM position_info WHERE position LIKE ? AND salary_low >= ? AND salary_high <= ?`
	COMPANY_SQL = `SELECT company_name, COUNT(*) AS num FROM position_info WHERE position LIKE ? GROUP BY company_name ORDER BY num DESC`
	WORKYEAR_SQL = `SELECT COUNT(*) AS num FROM position_info WHERE position LIKE ? AND level = ? GROUP BY workyear`
	EDUCAITON_SQL = `SELECT education, COUNT(*) AS num FROM position_info WHERE position LIKE ? GROUP BY education`
	LANGUAGE_SQL = `SELECT COUNT(*) AS num FROM position_info WHERE position LIKE ?`
)

var (
	db *sql.DB; err error
)

type PositionModel struct {
	Position string
}

func InitDB() {
	username := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpwd")
	protocol := beego.AppConfig.String("mysqlprorocol")
	address  := beego.AppConfig.String("mysqladdre")
	DBname   := beego.AppConfig.String("mysqlDBname")
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8", username, password, protocol, address, DBname))
	if err != nil {
		panic(err)
	}
}

func (p *PositionModel) GetCity(language string) *conf.Response {
	res := &conf.Response{}
	row, err := db.Query(CITY_SQL, "%" + language + "%")
	if err != nil {
		res.Code = 403
		res.Msg = "query error, " + err.Error()
		return res
	}
	res.Code = 200
	res.Msg = "success"
	for row.Next() {
		var city string
		var num int
		row.Scan(&city, &num)
		r := conf.NameToNum{city, num}
		res.Data = append(res.Data, r)
	}
	return res
}
// 5-10k, 10-15k, 15-20k, 20-25k, 25-30k, 35-40k, 40-45k, 45-50k, 50k以上
func (p *PositionModel) GetSalary(language string) *conf.Response {
	res := &conf.Response{}
	low := 5000
	high := 10000
	// 5-10k
	for ; low <= 50000; low += 5000{
		var num int
		data := conf.NameToNum{}
		high += 5000
		data.Name = fmt.Sprintf("%dk-%dk", low/1000, high/1000)
		if low == 50000 {
			high = low * 2
			data.Name = "50k以上"
		}
		row := db.QueryRow(SALARY_SQL, "%" + language + "%", low, high)
		row.Scan(&num)
		data.Num = num
		res.Data = append(res.Data, data)
	}
	res.Code = 200
	res.Msg = "success"
	return res
}

func (p *PositionModel) GetCompany(language string) *conf.Response {
	res := &conf.Response{}
	row, err := db.Query(COMPANY_SQL, "%" + language + "%")
	if err != nil {
		res.Code = 403
		res.Msg = "query error, " + err.Error()
		return res
	}
	res.Code = 200
	res.Msg = "success"
	cnt := 10
	for row.Next() {
		if cnt < 0 {
			break
		}
		data := conf.NameToNum{}
		row.Scan(&data.Name, &data.Num)
		res.Data = append(res.Data, data)
		cnt--
	}
	return res
}

func (p *PositionModel) GetWorkYear(language string) *conf.Response {
	res := &conf.Response{}
	workyear := []string{"经验不限", "应届生", "1年以内", "1-3年", "3-5年", "5-10年", "10年以上"}
	for i := 0; i <= 6; i++ {
		data := conf.NameToNum{}
		row := db.QueryRow(WORKYEAR_SQL, "%" + language + "%", i)
		row.Scan(&data.Num)
		data.Name = workyear[i]
		res.Data = append(res.Data, data)
	}
	res.Code = 200
	res.Msg = "success"
	return res
}

func (p *PositionModel) GetEducation(language string) *conf.Response {
	res := &conf.Response{}
	row, err := db.Query(EDUCAITON_SQL, "%" + language + "%")
	if err != nil {
		res.Code = 403
		res.Msg = "query error, " + err.Error()
		return res
	}
	for row.Next() {
		data := conf.NameToNum{}
		row.Scan(&data.Name, &data.Num)
		res.Data = append(res.Data, data)
	}
	res.Code = 200
	res.Msg = "success"
	return res
}

func (p *PositionModel) GetLanguage() *conf.Response {
	res := &conf.Response{}
	languageList := []string{"Golang", "PHP", "Node.js", "Java", "C++", "C#", "Python", "Ruby"}
	for _, v := range languageList {
		data := conf.NameToNum{}
		row := db.QueryRow(LANGUAGE_SQL, "%" + v + "%")
		row.Scan(&data.Num)
		data.Name = v
		res.Data = append(res.Data, data)
	}
	res.Code = 200
	res.Msg = "success"
	return res
}