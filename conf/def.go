package conf

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data []interface{} `json:"data"`
}

type NameToNum struct {
	Name string
	Num int
}

type DataList struct {
	Name string
	Data []NameToNum
}

type NameTo2Num struct {
	Name string
	Num1 int
	Num2 int
}

