package repository

type IStartRepo interface {
	Speak(message string) string
	Get(ID int)string
}
