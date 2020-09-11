package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"reflect"
)

type JsonReaderToJson struct{}

func (j JsonReaderToJson) IsAble(object interface{}) bool {
	jsonReader, ok := object.(inter.JsonReader)
	return ok && InterfaceToJson{}.IsAble(jsonReader.Json())
}

func (j JsonReaderToJson) EncodeThrough(object interface{}, encoders []inter.Encoder) (string, error) {
	jsonReader, ok := object.(inter.JsonReader)
	if !ok {
		return "", errors.New("can not transform to json with an unsupported type " + reflect.TypeOf(object).String())
	}

	return EncodeThrough(jsonReader.Json(), encoders)
}
