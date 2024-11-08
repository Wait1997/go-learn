package main

//import (
//	"go-learn/basic"
//)
//import "go-learn/func"

import (
	_interface "go-learn/interface"
	"go-learn/wait_group"
)

func main() {
	//basic.VarDeclareCase()
	//basic.StructCase()
	//basic.NewCase()
	//basic.MapInit1()
	//basic.MapInit2()
	//basic.MakeCase()
	//basic.SliceAndMapCase()
	//basic.ConvertCase()

	//user := _func.NewUser("nick", 18)
	//fmt.Println(user.GetName(), user.GetAge())

	//operate.FlowControlCase()

	//generic.SimpleCase()
	//generic.TTypeCase()
	//
	//ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	//defer stop()
	//<-ctx.Done()

	//generic.InterfaceCase()

	// standard.EncodingCase()
	// standard.ErrorsCase()
	// standard.FmtCase()
	// standard.FmtCase1()
	// standard.RegexpCase()

	//deferRecoverPanic.DeferCase1()
	//deferRecoverPanic.DeferCase2()

	//cat := _interface.NewCat()
	//animalLife(cat)
	//dog := _interface.NewDog()
	//animalLife(dog)

	//channel.Communication()
	//channel.ConcurrentSync()
	//
	//ch := make(chan os.Signal)
	//signal.Notify(ch, os.Interrupt, os.Kill)
	//<-ch

	//wait_group.WaitGroupCase()
	//wait_group.WaitGroupCase1()
	//ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	//defer stop()
	//<-ctx.Done()

	wait_group.CondCase()
}

func animalLife(a _interface.Animal) {
	a.Each()
	a.Drink()
	a.Sleep()
	a.Run()
}