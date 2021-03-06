package pack

import (
	"reflect"
	"fmt"
	"encoding/json"
)

type pointer struct {
	addr reflect.Value
	ttype reflect.Type

}

func (p *pointer) Set(val interface{}) {
	po := *p
	switch vv := val.(type) {
	case string:
		po.addr.SetString(vv)
	case int:
		po.addr.SetInt(int64(vv))
	case float64:
		switch po.addr.Kind() {
		default:
			po.addr.SetFloat(vv)
		case reflect.Int64,reflect.Int32,reflect.Int16,reflect.Int8,reflect.Int:
			po.addr.SetInt(int64(vv))
		}
	case bool:
		po.addr.SetBool(vv)
	case []interface{}:

	default:
		fmt.Println("todo feature")
	}
}

func Packer(b *[]byte, class interface{}) {

	v := reflect.ValueOf(class)

	indirect_type := reflect.Indirect(v).Type()

	//store addr
	name_addr := make(map[string]pointer)
	el := v.Elem()
	for i := 0; i < el.NumField(); i++ {
		name_addr[indirect_type.Field(i).Name] =
			pointer{v.Elem().Field(i), indirect_type}
	}


	var f interface{}
	json.Unmarshal(*b, &f)
	m := f.(map[string]interface{})

	for k, v := range m {
		if q, ok := name_addr[k]; ok {
			p := &q
			p.Set(v)
		}
	}
}
