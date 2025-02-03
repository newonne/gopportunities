package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string //para qual vaga,  ex: desenvolvedor backend
	Company  string //empresa que estar contratando
	Location string //  local da vaga
	Remote   bool   // se vaga remota ou nao
	Link     string // link da vaga
	Salary   int64  // salario
}
 