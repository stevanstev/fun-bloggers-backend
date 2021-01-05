package model

import (
	"net/http"
)

type BaseModel interface{
	FromJSON(r *http.Request) error
	ToJSON() string
}