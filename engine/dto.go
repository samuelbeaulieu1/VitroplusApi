package engine

import (
	"bytes"
	"encoding/json"
)

func ParseModelToDTO(model interface{}, dto interface{}) {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(model)
	json.NewDecoder(buffer).Decode(&dto)
}
