package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func PrettyJson(v interface{}) ([]byte, error) {
	indentNum := 4
	return json.MarshalIndent(v, "", strings.Repeat(" ", indentNum))
}

func ForceWriteFile(filename string, data []byte) error {
	var filePerm os.FileMode = 0644
	var dirPerm os.FileMode = 0755

	err := os.WriteFile(filename, data, filePerm)
	if err != nil {
		dir := filepath.Dir(filename)
		err := os.MkdirAll(dir, dirPerm)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, data, filePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func SliceContains(slice interface{}, element interface{}) bool {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		itemType := reflect.TypeOf(item)
		elementType := reflect.TypeOf(element)

		if !elementType.ConvertibleTo(itemType) {
			continue
		}

		elementValue := reflect.ValueOf(element)
		target := elementValue.Convert(itemType).Interface()

		if ok := reflect.DeepEqual(item, target); ok {
			return true
		}
	}
	return false
}
