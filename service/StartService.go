package service

import "github.com/bingjian-zhu/gin-inject/repository"

//StartService 注入IStartRepo
type StartService struct {
	Repo repository.IStartRepo `inject:""`
}

//Say 实现Say方法
func (s *StartService) Say(message string) string {
	return s.Repo.Speak(message)
}
