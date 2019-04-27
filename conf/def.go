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

