package main

import "fmt"

func updateName(name string)string{
	name="batuchi"
	return name
}

func updateName2(name *string){
	*name="bear"
}

func main(){
	menu:=map[string]float64{
		"soup":4.6,
		"cake":3.5,
		"oil":4.4,
	}
	fmt.Println(menu["cake"])

	//loopin maps
	for k,v:=range menu{
		fmt.Println(k,"-",v)
	}
	delete(menu,"oil")
	for k,v:=range menu{
		fmt.Println(k,"-",v)
	}

	name:="shaharin"
	fmt.Println("Name before update",name)
	name=updateName("shaharin")
	fmt.Println("Name after update",name)
	namePointer:=&name
	fmt.Println(*namePointer)
	fmt.Println("Name before update",name)
	updateName2(namePointer)
	fmt.Println("Name after update",name)


	//struct
	mybill:=newBill("shaharin")
	fmt.Println(mybill)

}