package deepcopy

import (
	"fmt"
	"github.com/json-iterator/go"
	"reflect"
)

func Copy(dst interface{}, src interface{}) error {
	if dst == nil {
		return fmt.Errorf("dst cannot be nil")
	}
	if src == nil {
		return fmt.Errorf("src cannot be nil")
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("Unable to marshal src: %s", err)
	}
	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal into dst: %s", err)
	}
	return nil
}

func Clone(src interface{}) (interface{}, error) {
	if src == nil {
		return nil, fmt.Errorf("src cannot be nil")
	}
	t := reflect.TypeOf(src)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	dstValue := reflect.New(t)
	err := Copy(dstValue.Interface(), src)
	return dstValue.Interface(), err
}
