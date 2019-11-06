package repository

import (
	"fmt"
	"github.com/bingjian-zhu/gin-inject/common/datasource"
)

//StartRepo 注入数据库
type StartRepo struct {
	Source datasource.IDb `inject:""`
}

//Speak 实现Speak方法
func (s *StartRepo) Speak(message string) string {
	return fmt.Sprintf("[Repository] speak: %s", message)
}
