package main

import "fmt"

func main(){

	//string
	var nameOne string="shaharin"
	var nameTwo="ahmed"
	var nameThree string
	nameThree="refat"

	//short hand for var nameFour string="omar" cant use this shorthand outside of function
	nameFour:="omar"
	fmt.Println(nameOne,nameTwo,nameThree,nameFour)


	//int
	var numberOne int=32
	var numberTwo=33
	var numberThree int
	numberThree=34
	numberFour:=35
	// var numberFive int8=1
	fmt.Println(numberOne,numberTwo,numberThree,numberFour)


	//print
	fmt.Print("My name is ",nameOne," my age is ",numberOne,"\n")
	fmt.Printf("my name is %v which is of type %T and my age is %v which is of type %T\n",nameOne,nameOne,numberOne,numberOne)

	//Sprintf save formatted strings
	var str=fmt.Sprintf("My name is %v and my age is %v \n",nameOne,numberOne)
	println("Saved string is",str)

	//arrays
	var ages=[4]int{1,3,4,5}
	fmt.Println(ages,len(ages))

	var names=[3]string{"shaharin","ahmed","refat"}
	fmt.Println(names)
	fmt.Println(names[2])

	var arr1=[...]int{1,3,4}
	fmt.Println(arr1)

	arr2:=[...]int{3,5,4}
	fmt.Println(arr2)

	var arr3 =[2]int{}
	arr3[0]=4
	arr3[1]=5
	fmt.Println(arr3)

	arr4 := [5]int{3:10,0:40}
	fmt.Println(arr4)

	//slices(uses array under the hood)
	ages1:=ages[0:3]
	ages2:=ages[:3]
	ages3:=ages[2:]
	fmt.Println(ages1)
	fmt.Println(ages2)
	fmt.Println(ages3)
	ages1=append(ages1, 45);
	fmt.Println(ages1)
	ages1=append(ages1, ages3...)
	fmt.Println(ages1)
	ages1=append(ages1, 69,96)
	fmt.Println(ages1)

	//difference between len and cap in slices

	
}