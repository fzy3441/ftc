package main

import (
	"fmt"
	"github.com/fzy3441/ftc/ch"
	"github.com/fzy3441/ftc/tt"
	"time"
)

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
	tool.Wait()*/
	//select{}
	//
	//link := ch.NewLink()
	//tt.GoFor(func(o *tt.Gfo) {
	//	go func(key int,link *ch.Link) {
	//		for i:=0;i<5;i++ {
	//			link.Push(fmt.Sprintf("k=>%+v v=>%+v===>",key,i))
	//			time.Sleep(1*time.Second)
	//		}
	//	}(o.Num,link)
	//
	//	o.SetBreak(o.Num >=50)
	//}).Action()
	//time.Sleep(5*time.Second)
	//fmt.Printf("count=>%+v\n",link.Count)
	//
	//link.Range(func(o *tt.Gfo, node *ch.Node) {
	//	fmt.Printf("-------->>>%+v====>>>>%+v\n",o,node)
	//},ch.RANGE_STATUS_ORDER_DESC)

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

	//fmt.Printf("%+v",link.List())
	//fmt.Printf("%+v",link)
	//
	//link.Insert(7,10)
	//fmt.Printf("%+v",link)
	//fmt.Printf("%+v",link.List())
	//
	//link.Insert(0,11)
	//fmt.Printf("%+v",link.List())
	//fmt.Println(link.Count)
	//fmt.Printf("%+v",link.List())

	//fmt.Println(tt.Ternary(false==true,1,aa))

/*	fmt.Println(tt.TimeFmt())
*/
	//link := ch.NewLink()
	//tool:=tt.NewTool(5)
	//tool.Run()
	//
	//for  i:=0;i<5;i++{
	//	tool.AddTask(func(w *tt.Worker) {
	//		link.Push(1)
	//		fmt.Printf("name=>%s  num=>%d i=>>%d\n",w.Name,222,i)
	//		time.Sleep(1*time.Second)
	//
	//	})
	//}
	//tool.Wait()
	//fmt.Println("===================>")
	//for  i:=0;i<5;i++{
	//	tool.AddTask(func(w *tt.Worker) {
	//		link.Push(1)
	//		fmt.Printf("name=>%s  num=>%d i=>>%d\n",w.Name,333,i)
	//		time.Sleep(1*time.Second)
	//
	//	})
	//}
	//tool.Wait()
	//fmt.Printf("link count=>>%d\n",link.Count)


	//cache :=ch.NewCacher()
	a:=ch.NewLink()
	a.Push(1)
	a.Push(2)
	a.Push(3)
	a.Push(4)
	//b:=ch.NewLink()
	//b.Push(1)
	//b.Push(2)
	//b.Push(3)
	//b.Push(4)
	//a.Merge(b)
	//a.RemoveAt(2)
	//a.Remove(a.Head)
	//a.Remove(a.Head)
	//a.Remove(a.Head)
	//a.Remove(a.Head)
	//a.Remove(a.Head.Next.Next)
	go func() {
		a.Range(func(o *tt.Gfo, node *ch.Node) {
			fmt.Printf("----------%d\n",node.Data.(int))
			fmt.Printf("==========%+v\n",a.Clone().List())

			time.Sleep(2*time.Second)
		})
	}()



	go func() {
		a.Range(func(o *tt.Gfo, node *ch.Node) {
			fmt.Printf("----------%d\n",node.Data.(int))
			fmt.Printf("==========%+v\n",a.Clone().List())

			time.Sleep(3*time.Second)
		})
	}()

	fmt.Println(a.Count)
	time.Sleep(10*time.Second)



}
