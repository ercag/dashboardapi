package models

type ResponseModel struct {
	ResCode    int
	ResMessage string
	ResData    interface{}
}
