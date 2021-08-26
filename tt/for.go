package tt

import (
	"fmt"
	"sync"
)

type FORSTATUS_TYPE_INT int

const (
	_FORSTATUS_CONNETION FORSTATUS_TYPE_INT = iota
	_FORSTATUS_STOP
	_FORSTATUS_PAUSE
)

type FStus struct {
	_wg *sync.WaitGroup
	_cmd    chan FORSTATUS_TYPE_INT
	_status   FORSTATUS_TYPE_INT
	_exec     func(o *Gfo)
	_cond   func(o *Gfo)bool
	Option *Gfo
}

type Gfo struct {
	_break bool
	Num , Inx int
	Cache interface{}
}

func (obj *Gfo)Break()  {
	obj._break=true
}

func (obj *FStus)SetWaitGroup(wg *sync.WaitGroup)  {
	obj._wg=wg
}

func GoFor(f func(o *Gfo))*FStus {
	fs:=&FStus{
		_cmd: make(chan FORSTATUS_TYPE_INT),
		_status: _FORSTATUS_CONNETION,
		_cond: func(o *Gfo) bool {
			return true
		},
		_exec: f,
		Option: &Gfo{
			Num: 0,
			Inx: 0,
		},
	}
	return fs
}
func (obj *FStus)WaitGroup(wg *sync.WaitGroup)  {
	obj._wg=wg

}

func (obj *FStus)Cond(f func(o *Gfo)bool)*FStus  {
	obj._cond = f
	return obj
}

func (obj *FStus)Exec(f func(o *Gfo))*FStus  {
	obj._exec = f
	return obj
}

func (obj *FStus)Action()  {
	if obj._wg != nil {
		fmt.Println("++")
		obj._wg.Add(1)
	}
	for obj.checkgo() {
		select {
		case status, ok := <-obj._cmd:
			if ok {
				obj._status = status
			}
		default:
		}

		switch obj._status {
		case _FORSTATUS_CONNETION:
			obj._exec(obj.Option)
			obj.Option.Num++
			obj.Option.Inx++
		case _FORSTATUS_STOP:
			return
		case _FORSTATUS_PAUSE:
			continue
		}
	}
	if obj._wg != nil {
		obj._wg.Done()
	}
}
func (obj *FStus)Run() {
		go obj.Action()
}

func (obj *FStus)checkgo()bool  {
	return obj._cond(obj.Option)&&!obj.Option._break
}

func (obj *FStus)Pause()  {
	fmt.Println("pause")
	obj._cmd <- _FORSTATUS_PAUSE
}

func (obj *FStus)Stop()  {
	obj._cmd <- _FORSTATUS_STOP
}

func (obj *FStus)Continue()  {
	fmt.Println("continue")
	obj._cmd <-_FORSTATUS_CONNETION
}
