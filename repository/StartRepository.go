package repository

import (
	"fmt"
	"gin-inject/common/datasource"
)

type StartRepo struct {
	Source datasource.IDb `inject:""`
}

func (s *StartRepo) Speak(message string) string {
	return fmt.Sprintf("[Repository] speak: %s", message)
}
