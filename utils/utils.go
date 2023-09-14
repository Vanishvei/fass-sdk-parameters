package utils

// File       : utils.go
// Path       : utils
// Time       : CST 2023/6/6 18:25
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

func Bool(a bool) *bool {
	return &a
}

func BoolValue(a *bool) bool {
	if a == nil {
		return false
	}
	return *a
}

func String(a string) *string {
	return &a
}

func StringValue(a *string) string {
	if a == nil {
		return ""
	}
	return *a
}

func Int(a int) *int {
	return &a
}

func IntValue(a *int) int {
	if a == nil {
		return 0
	}
	return *a
}

func StringSlice(a []string) []*string {
	if a == nil {
		return nil
	}
	res := make([]*string, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func StringSliceValue(a []*string) []string {
	if a == nil {
		return nil
	}
	res := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Map(a map[string]string) *map[string]string {
	if len(a) == 0 {
		return nil
	}
	res := make(map[string]string)
	for k, v := range a {
		res[k] = v
	}
	return &res
}

func MapValue(a *map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	res := make(map[string]string)
	for k, v := range *a {
		res[k] = v
	}
	if len(res) == 0 {
		return nil
	}
	return res
}

func DefaultNumber(reaNum, defaultNum *int) *int {
	if reaNum == nil {
		return defaultNum
	}
	return reaNum
}
