package repository

//IStartRepo 定义IStartRepo接口
type IStartRepo interface {
	Speak(message string) string
}
