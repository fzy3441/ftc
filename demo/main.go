package main

func main() {
/*	fmt.Println("hello word")
	f:=tt.GoFor(func(o *tt.Gfo){
		fmt.Printf("num==%d\n",o.Num)
		time.Sleep(1*time.Second)
	})
	f.Run()
	time.Sleep(3*time.Second)
	f.Pause()
	time.Sleep(4*time.Second)
	f.Continue()
	time.Sleep(4*time.Second)

	f.Stop()
	fmt.Println("over")*/

	/*tt.Try(func() {
		panic("hello world")
	}).Cache(errors.New("hello world"), func(status *tt.TryStatus) {
		fmt.Println("aaaaaaaaaaa")
	}).Cache(errors.New("hello world1"), func(status *tt.TryStatus) {
		fmt.Println("aaaaaaaaaaa1")
	}).Cache(errors.New("hello world2"), func(status *tt.TryStatus) {
		fmt.Println("aaaaaaaaaaa2")
	}).Cache(errors.New("hello world3"), func(status *tt.TryStatus) {
		fmt.Println("aaaaaaaaaaa3")
	}).Cache(errors.New("hello world4"), func(status *tt.TryStatus) {
		fmt.Println("aaaaaaaaaaa4")
	}).Finally(func() {
		fmt.Println("finally")
	},func() {
			fmt.Println("finally1")
		},func() {
				fmt.Println("finally2")
		},func() {
					fmt.Println("finally3")
		},func() {
			fmt.Println("finally4")
		})*/

	/*tool:=tt.NewTool(2)
	tool.Run()
	tool.AddTask(func(w *tt.Worker) {
		tt.GoFor(func(o *tt.Gfo) {
			fmt.Println(w.Name,"111")
			time.Sleep(time.Second*1)
		}).Cond(func(o *tt.Gfo) bool {
			return o.Num <5
		}).Action()

	})
	tool.AddTask(func(w *tt.Worker) {
		tt.GoFor(func(o *tt.Gfo) {
			fmt.Println(w.Name,"222")
			time.Sleep(time.Second*1)
		}).Cond(func(o *tt.Gfo) bool {
			return o.Num <5
		}).Action()

	})
	tool.AddTask(func(w *tt.Worker) {
		tt.GoFor(func(o *tt.Gfo) {
			fmt.Println(w.Name,"333")
			time.Sleep(time.Second*1)
		}).Cond(func(o *tt.Gfo) bool {
			return o.Num <5
		}).Action()

	})
	select{}*/

/*	link := ch.NewLink()
	link.Push(1)
	link.Push(2)
	link.Push(3)
	link.Push(4)
	link.Push(5)
	link.Push(6)
	link.Push(7)
	link.Push(8)
	link.Push(9)
	fmt.Printf("%+v",link.List())
	link.Delete(3)
	fmt.Printf("%+v",link.List())*/

	//fmt.Printf("%+v\n",link.Pop())
	//fmt.Printf("%+v\n",link.Pop())
	//fmt.Printf("%+v\n",link.Pop())
	//fmt.Printf("%+v\n",link.Pop())
	//fmt.Printf("%+v\n",link.Pop())
	//fmt.Printf("%+v\n",link.Pop())
	//fmt.Printf("%+v\n",link.Pop())
	//fmt.Printf("%+v\n",link.Pop())
	//fmt.Printf("%+v\n",link.Pop())
	//fmt.Printf("%+v\n",link.Pop())

/*	fmt.Printf("%+v",link.List())
	fmt.Printf("%+v",link)

	link.Insert(8,10)
	fmt.Printf("%+v",link)
	fmt.Printf("%+v",link.List())

	link.Insert(0,11)
	fmt.Printf("%+v",link.List())
	fmt.Println(link.Count)
	fmt.Printf("%+v",link.List())*/

/*	fmt.Println(tt.TimeFmt())
*/

	
}
