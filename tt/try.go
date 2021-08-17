package tt

import (
	"fmt"
	"reflect"
)

type TryStatus struct {
	_cacheFlag bool
	Err error
}

func Try(fn func())*TryStatus {
	ts:=&TryStatus{}
	defer func()*TryStatus{
		if err := recover(); err != nil{
			ts.Err = fmt.Errorf("%v", err)
		}
		return ts
	}()
	defer
	fn()
	return ts
}

func (obj *TryStatus)checkCacheFlag()bool  {
	return obj._cacheFlag
}

func (obj *TryStatus)Cache(err error,fn func(status *TryStatus))*TryStatus  {
	if obj.checkCacheFlag() {
		return obj
	}

	if reflect.TypeOf(err) == reflect.TypeOf(obj.Err) {
		fn(obj)
		obj._cacheFlag= true
	}

	return obj
}

func (obj *TryStatus)Finally(fns ...func())  {
	for  _,fn := range fns{
		defer fn()
	}
}
