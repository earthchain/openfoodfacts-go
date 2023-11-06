package openfoodfacts

import (
	"fmt"
	"strconv"
)

func AsFloat64(val interface{}) *float64 {

	switch val.(type) {
	case nil:
		return nil
	case int:
	case int16:
	case int32:
	case int64:
	case float32:
	case float64:
		fv := val.(float64)
		return &fv
	case string:
		fv, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return nil
		}
		return &fv
	}
	return nil
}

func AsString(val interface{}) *string {

	switch val.(type) {
	case nil:
		return nil
	case int:
		fv := fmt.Sprintf("%d", val.(int))
		return &fv
	case int16:
		fv := fmt.Sprintf("%d", val.(int16))
		return &fv
	case int32:
		fv := fmt.Sprintf("%d", val.(int32))
		return &fv
	case int64:
		fv := fmt.Sprintf("%d", val.(int64))
		return &fv
	case float32:
		fv := fmt.Sprintf("%f", val.(float32))
		return &fv
	case float64:
		fv := fmt.Sprintf("%f", val.(float64))
		return &fv
	case string:
		fv := val.(string)
		return &fv
	}
	return nil
}
