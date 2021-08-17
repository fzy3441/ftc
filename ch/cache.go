package ch

type Cacher struct {
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
	obj._cmap[key]=data
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
	return obj.List()
}

func (obj *Cacher)Maps() map[string]interface{} {
	return obj._cmap
}
