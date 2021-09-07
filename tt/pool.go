package tt

import (
	"time"
)

type StatStopInt int

const (
	StaStopt StatStopInt = 0
	GEN_RAND_STR_LEN = 12
	CHAN_QUEUE_LEN_TAKE = 10

)

type Stat struct {
	iWorker,iWorking,iIdle,iQueue,iCount int
	chQueueI,chIdleI,chWorkingI chan int
	chComm chan StatStopInt
}

func newStat(num int)*Stat  {
	stat:= &Stat{
		iWorker: num,
		chIdleI: make(chan int,0),
		chQueueI: make(chan int,0),
		chWorkingI: make(chan int,0),
		chComm: make(chan StatStopInt,0),
	}
	go stat.Run()
	return stat
}

func (obj *Stat)Run()  {
	for {
		select {
		case i := <-obj.chIdleI:
			obj.iIdle += i
			//fmt.Printf("=-=09099099-=-=-%+v\n",obj)
		case i := <-obj.chQueueI:
			obj.iQueue += i
		case i := <-obj.chWorkingI:
			obj.iWorking += i
		case <-obj.chComm:
			return
		}
	}
}

func (obj *Stat) ResetCount()  {
	obj.iCount=0
}

func (obj *Stat)Count()int  {
	return obj.iCount
}

type Pool struct {
	*Stat
	chQueueFn   chan func(w *Worker)
	chIdleWork chan *Worker
	mpWorking map[string] *Worker

}

type Worker struct {
	Name string
	Task func(w *Worker)
}

func newWorker() *Worker {
	return &Worker{
		Name: GenRandStr(GEN_RAND_STR_LEN),
	}
}
func (obj *Worker)Run()  {
	obj.Task(obj )
}

func NewTool(num int)*Pool  {
	pool:= &Pool{
		Stat:newStat(num),
		chQueueFn: make(chan func(w *Worker),num*CHAN_QUEUE_LEN_TAKE),
		mpWorking: make(map[string]*Worker),
		chIdleWork: make(chan *Worker,num),
	}
	pool.init()
	return pool
}
func (obj *Pool)init()  {
	obj.initWorker()
}

func (obj *Pool)actionWorking()  {
	GoFor(func(o *Gfo) {
		work := <-obj.chIdleWork
		obj.chIdleI <- -1
		work.Task = <-obj.chQueueFn
		obj.chQueueI <- -1

		obj.newWork(work)
	}).Run()
}

func (obj *Pool)Wait()  {
	for  {
		if obj.iQueue==0&&obj.iWorking==0{
			return
		}
		time.Sleep(1*time.Second)
	}
}

func (obj *Pool)AddTask(task func(w *Worker))  {
	obj.chQueueFn<- task
	obj.chQueueI <- 1
}

func (obj *Pool) initWorker(){
	GoFor(func(o *Gfo) {
		obj.chIdleWork <- newWorker()
		obj.chIdleI <- +1
		o.SetBreak(o.Num+1>=obj.iWorker)
	}).Run()
}
func (obj *Pool)newWork(work *Worker)  {
	go func(work *Worker) {
		obj.chWorkingI<- +1
		work.Run()
		obj.chWorkingI<- -1
		obj.iCount++
		obj.chIdleWork <- work
		obj.chIdleI <- +1
	}(work)
}
func (obj *Pool)Run()  {
	obj.ResetCount()
	obj.actionWorking()
}

func (obj *Pool)Stop()  {
	obj.chComm<- StaStopt
	return
}


