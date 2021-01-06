package model

import (
	"net/http"
)

/*BaseModel ...
@desc interface to set mandatory functions to it's child
*/
type BaseModel interface {
	FromJSON(r *http.Request) error
	ToJSON() string
}
