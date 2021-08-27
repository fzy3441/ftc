package ch

import (
	"sync"
)

type Cacher struct {
	mutex sync.Mutex
	_cmap map[string]interface{}
	_link_list *Link
}

func NewCacher() *Cacher {
	return &Cacher{
		_cmap: make(map[string]interface{}),
		_link_list: NewLink(),
	}
}

func (obj *Cacher)Set(key string,data interface{}){
	obj.mutex.Lock()
	obj._cmap[key]=data
	obj.mutex.Unlock()
}

func (obj *Cacher)Get(key string)interface{}{
	return obj._cmap[key]
}

func (obj *Cacher)Push(data interface{})  {
	obj._link_list.Push(data)
}

func (obj *Cacher)Pop()interface{} {
	return obj._link_list.Pop()
}

func (obj *Cacher)List() []interface{} {
	return obj._link_list.List()
}

func (obj *Cacher)Maps() map[string]interface{} {
	return obj._cmap
}

func (obj *Cacher)IsExists(key string)bool  {
	_,ok :=obj._cmap[key]
	return ok
}

type ForStatus struct {
	_break bool
	Inx int
}

func (obj *ForStatus)SetBreak(flag bool)  {
	obj._break = flag
}

func (obj *ForStatus)Break()  {
	obj.SetBreak(true)
}

type Value struct {
	k string
	v interface{}
}

func (obj *Cacher)RangeMap(f func(fs *ForStatus,val *Value))  {
	fs:=&ForStatus{}
	for k,v := range obj._cmap{
		fs.Inx++
		f(fs,&Value{
			k: k,
			v: v,
		})
		if fs._break {
			break
		}
	}
}