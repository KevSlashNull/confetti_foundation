package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"reflect"
)

func EncodeThrough(object interface{}, encoders []inter.Encoder) (string, error) {
	for _, encoder := range encoders {
		if encoder.IsAble(object) {
			return encoder.EncodeThrough(object, encoders)
		}
	}

	if err, ok := object.(error); ok {
		err := errors.New("No encoder found to handle error: " + err.Error())
		return err.Error(), err
	}

	err := errors.New("no encoder found to encode response body with type " + reflect.TypeOf(object).String())
	return EncodeThrough(err, encoders)
}
