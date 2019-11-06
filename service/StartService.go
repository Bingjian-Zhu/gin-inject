package service

import "gin-inject/repository"

type StartService struct {
	Repo repository.IStartRepo `inject:""`
}

func (s *StartService) Say(message string) string {
	return s.Repo.Speak(message)
}
