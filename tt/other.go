package tt

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

func GenRandStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func TestFnTime(f interface{}) string {
	start := time.Now()
	callFn(f)
	end := time.Now()
	vf := reflect.ValueOf(f)
	str := fmt.Sprintf("[%s] runtime: %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
	fmt.Println(str)
	return str
}

func callFn(f interface{}) interface{} {
	if f != nil {
		t := reflect.TypeOf(f)
		if t.Kind() == reflect.Func && t.NumIn() == 0 {
			function := reflect.ValueOf(f)
			in := make([]reflect.Value, 0)
			out := function.Call(in)
			if num := len(out); num > 0 {
				list := make([]interface{}, num)
				for i, value := range out {
					list[i] = value.Interface()
				}
				if num == 1 {
					return list[0]
				}
				return list
			}
			return nil
		}
	}
	return f
}

func isZero(f interface{}) bool  {
	v := reflect.ValueOf(f)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		str := v.String()
		if str == "" {
			return true
		}
		zero, err := strconv.ParseFloat(str, 10)
		if zero == 0 && err == nil {
			return true
		}
		boolean, err := strconv.ParseBool(str)
		return boolean == false && err == nil
	default:
		return false
	}
}


func If(args ...interface{}) interface{} {
	var cond = callFn(args[0])
	if len(args) == 1 {
		return cond
	}
	var trueVal = args[1]
	var falseVal interface{}
	if len(args) > 2 {
		falseVal = args[2]
	} else {
		falseVal = nil
	}
	if cond == nil {
		return callFn(falseVal)
	} else if v, ok := cond.(bool); ok {
		if v == false {
			return callFn(falseVal)
		}
	} else if isZero(cond) {
		return callFn(falseVal)
	} else if v, ok := cond.(error); ok {
		if v != nil {
			fmt.Println(v)
			return cond
		}
	}
	return callFn(trueVal)
}

func Or(args ...interface{}) interface{} {
	var cond = callFn(args[0])
	if len(args) == 1 {
		return cond
	}
	if cond == nil {
		return callFn(args[1])
	}
	if v, ok := cond.(bool); ok {
		if v == false {
			return callFn(args[1])
		}
	} else if isZero(cond) {
		return callFn(args[1])
	} else if v, ok := cond.(error); ok {
		if v != nil {
			fmt.Println(v)
			return cond
		}
	}
	return cond
}