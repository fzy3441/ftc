package tt

import "time"

type Pool struct {
	WorkNum int
	WorkingNum int
	IdleNum  int
	QueueNum int
	_queue   chan func(w *Worker)
	_working map[string] *Worker
	_idle chan *Worker
}

type Worker struct {
	Name string
	Task func(w *Worker)
}

func newWorker() *Worker {
	return &Worker{
		Name: GenRandStr(12),
	}
}
func (obj *Worker)Run()  {
	obj.Task(obj )
}

func NewTool(num int)*Pool  {
	return &Pool{
		WorkNum: num,
		_queue: make(chan func(w *Worker),num*2),
		_working: make(map[string]*Worker),
		_idle: make(chan *Worker,num),
	}
}

func (obj *Pool)working()  {
	gf:=GoFor(func(o *Gfo) {
		work := <-obj._idle
		obj.IdleNum--
		task := <-obj._queue
		obj.QueueNum--
		work.Task=task
		obj.newWork(work)
	})
	gf.Run()
}

func (obj *Pool)Wait()  {
	for  {
		if obj.QueueNum==0&&obj.WorkingNum==0{
			return
		}
		time.Sleep(1*time.Second)
	}
}

func (obj *Pool)AddTask(task func(w *Worker))  {
	obj.QueueNum++
	obj._queue<- task
}

func (obj *Pool) initWorker(){
	GoFor(func(o *Gfo) {
		obj._idle <- newWorker()
		obj.IdleNum++
	}).Cond(func(o *Gfo) bool {
		return o.Num <obj.WorkNum
	}).Action()
}
func (obj *Pool)newWork(work *Worker)  {
	go func() {
		obj.WorkingNum++
		work.Run()
		obj.WorkingNum--
		obj._idle <- work
		obj.IdleNum++
	}()
}
func (obj *Pool)Run()  {
	obj.initWorker()
	obj.working()
}


