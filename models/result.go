package models

type Result struct {
	Error string      `json:"error"`
	Data  interface{} `json:"date"`
	Code  int         `json:"code"`
}
