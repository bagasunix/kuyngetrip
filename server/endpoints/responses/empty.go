package responses

import (
	"encoding/json"

	"github.com/bagasunix/kuyngetrip/pkg/err"
)

type Empty struct{}

func (a *Empty) ToJSON() []byte {
	j, errs := json.Marshal(a)
	err.HandlerReturnedVoid(errs)
	return j
}
