package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}

// o uso de "json" nessa estrutura serve para informar ao Go que, quando
// um objeto do tipo OpeningResponse for convertido para json (atraves do pkg encoding/json),
// os campos devem ser nomeados com os nomes especificados na tag "json".

// Por exemplo, quando um objeto do tipo OpeningResponse for convertido para json,
// o campo "ID" do objeto sera nomeado como "id" no json gerado.

// Esse recurso eh util para quando queremos que o nome do campo no json seja diferente
// do nome do campo no Go. Por exemplo, se quisermos que o campo "CreatedAt" seja nomeado
// como "created_at" no json, podemos usar a tag "json" para informar ao Go sobre essa convencao.

// Alem disso, tambem podemos usar a tag "json" para informar ao Go que um campo deve ser omitido
// quando um objeto for convertido para json. Por exemplo, se quisermos que o campo "DeletedAt" seja
// omitido quando um objeto do tipo OpeningResponse for convertido para json, podemos usar a tag
// "json" para informar ao Go sobre essa convencao. Nesse caso, o campo "DeletedAt" sera omitido
// do json gerado.
