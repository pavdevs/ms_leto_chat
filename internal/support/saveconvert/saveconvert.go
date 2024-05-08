package saveconvert

import (
	"fmt"
	"strconv"
)

func SafeConvertToInt64(val any) (int64, error) {
	switch v := val.(type) {
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case float64:
		return int64(v), nil // Будьте осторожны, так как это может привести к потере точности
	case string:
		// Преобразование строки в int64
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("unsupported type %T", val)
	}
}
