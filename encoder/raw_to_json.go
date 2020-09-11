package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"reflect"
)

type RawToJson struct{}

func (v RawToJson) IsAble(object interface{}) bool {
	value, ok := object.(interface{ Raw() interface{} })
	return ok && InterfaceToJson{}.IsAble(value.Raw())
}

func (v RawToJson) EncodeThrough(object interface{}, encoders []inter.Encoder) (string, error) {
	value, ok := object.(interface{ Raw() interface{} })
	if !ok {
		return "", errors.New("can not transform to json with an unsupported type " + reflect.TypeOf(object).String())
	}

	result := value.Raw()

	// If the object is nil, we don't want to return an empty body
	if result == nil {
		return "null", nil
	}

	return EncodeThrough(result, encoders)
}
