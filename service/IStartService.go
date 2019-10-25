package service

type IStartService interface {
	Say(message string) string
	GetID(ID int)string
}
