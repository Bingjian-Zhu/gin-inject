package repository

import (
	"gin-inject/common/datasource"
	"fmt"
)

type StartRepo struct {
	Source datasource.IDb `inject:""`
}

func (s *StartRepo) Speak(message string) string {
	return fmt.Sprintf("[Repository] speak: %s", message)
}

func (s *StartRepo) Get(ID int)string{
	return fmt.Sprintf("[You] get: %d", ID)
}
