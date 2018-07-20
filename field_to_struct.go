package frame

import (
	"fmt"
	"reflect"
	"strings"
)
/**
 * FieldToTruct 数据库表字段与结构对应， sql查询结果为struct
 * @param {[type]} st interface{}) (string, []interface{} [description]
 */
func FieldToStruct(st interface{}) (string, []interface{}) {
	rStruct := reflect.ValueOf(st).Elem()
	rType := rStruct.Type()

	values := make([]interface{}, rType.NumField())
	fields := make([]string, rType.NumField())

	for i, rlen := 0, rType.NumField(); i < rlen; i++ {
		if rType.Kind() == reflect.Struct {
			fields[i] = rType.Field(i).Tag.Get("sql")
			values[i] = rStruct.FieldByName(rType.Field(i).Name).Addr().Interface()
		}
	}

	return fmt.Sprintf("`%s`", strings.Join(fields, "`,`")), values
}
