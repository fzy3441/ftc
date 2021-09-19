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
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	return obj._cmap[key]
}

func (obj *Cacher)GetVal(key string,data interface{})interface{}  {
	if !obj.IsExists(key){
		obj.Set(key,data)
	}
	return obj.Get(key)
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

func (obj *Cacher)GetLink()*Link  {
	return obj._link_list
}

func (obj *Cacher)IsExists(key string)bool  {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()
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
	Key string
	Value interface{}
}

func (obj *Cacher)RangeMap(f func(fs *ForStatus,val *Value))  {
	fs:=&ForStatus{}
	for k,v := range obj._cmap{
		fs.Inx++
		f(fs,&Value{
			Key: k,
			Value: v,
		})
		if fs._break {
			break
		}
	}
}