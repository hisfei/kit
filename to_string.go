/**
 * @Author: fei
 * @Description:
 * @File:  to_string
 * @Version: 1.0.0
 * @Date: 2023/11/16 18:28
 */
package kit

import (
	"encoding/json"
	"strconv"
	"time"
)

func ForceString(value interface{}) string {
	switch value.(type) {
	case time.Time:
		return value.(time.Time).Format("2006-01-02 15:04:05")

	case int:
		return strconv.Itoa(value.(int))
	case uint:

		return strconv.Itoa(int(value.(uint)))
	case int8:

		return strconv.Itoa(int(value.(int8)))
	case uint8:

		return strconv.Itoa(int(value.(uint8)))
	case int16:

		return strconv.Itoa(int(value.(int16)))
	case uint16:

		return strconv.Itoa(int(value.(uint16)))
	case int32:

		return strconv.Itoa(int(value.(int32)))
	case uint32:

		return strconv.Itoa(int(value.(uint32)))
	case int64:

		return strconv.FormatInt(value.(int64), 10)
	case uint64:

		return strconv.FormatUint(value.(uint64), 10)
	case string:
		return value.(string)

	case []byte:
		return string(value.([]byte))
	case bool:
		return strconv.FormatBool(value.(bool))
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'f', -1, 64)

	default:
		newValue, _ := json.Marshal(value)
		return string(newValue)

	}

	return "" //[1 : len(res)-1]
}
func ToString(v interface{}) string {
	t := ForceString(v)
	if t != "[]interface{}" {
		return t
	}
	res := "["
	for _, vcc := range v.([]interface{}) {
		if res != "[" {
			res = res + ","
		}
		res = res + ToString(vcc)
	}
	return res + "]"
}
