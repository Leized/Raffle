package model

type Data struct {
	ID     int
	Member string
	Award  string
}
type JsonResult struct {
	Code int
	Msg  string
	Data []Data
}
