package response

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"tsqqqqqq/util"
)

type Impl struct {
	P interface{}
}

func (i *Impl) Insert() (res sql.Result) {
	t, v := returnElem(i.P)
	len := v.NumField()
	var field string
	var value string
	var tableName string
	tableName = t.Name()
	values := make([]interface{}, 0, len)
	sql := "insert into " + tableName + "("
	for i := 0; i < len; i++ {
		key := t.Field(i)
		val := v.FieldByName(key.Name)
		if val.Interface() != "" && val.Interface() != 0 {
			field += key.Tag.Get("field") + ","
			value += "?,"
			values = append(values, val.Interface())
		}
	}
	field = strings.TrimRight(field, ",")
	sql += field + ")  value(" + strings.TrimRight(value, ",") + ")"
	fmt.Println(sql, values)
	res, err := util.ExecData(sql, values)
	if err != nil {
		panic(err)
	}
	return res
}

func returnElem(obj interface{}) (t reflect.Type, v reflect.Value) {
	t = reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	v = reflect.ValueOf(obj)
	v = v.Elem()
	return t, v
}
