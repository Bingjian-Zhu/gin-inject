package repository

type IStartRepo interface {
	Speak(message string) string
}
