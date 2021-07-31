package main

import "fmt"

//package level variable of type map[]
var data map[int]Customer

type dat map[int]Customer

//Defined a struct named Customer
type Customer struct {
	Cid      int
	Cname    string
	Caddress string
	Cmob     uint64
}

var c Customer

//intialize map with make
func init() {
	data = make(map[int]Customer)

}

//receiver function for add
func (a Customer) add(k int, v Customer) {
	//check if key exist or not...
	if _, ok := data[k]; ok {
		fmt.Println("customer exists")
	} else {
		data[k] = v
	}
}

//receiver function for update
func (a Customer) update(k int, v Customer) {
	if _, ok := data[k]; ok {
		data[k] = v
		fmt.Println("your data is updated")
	} else {
		fmt.Println("not updated.....please enter valid id")
	}
}

//receiver function for delete
func (a Customer) delete(k int) {
	if _, ok := data[k]; ok {
		delete(data, k)
		fmt.Println("your data is deleted successfully.......")
	} else {
		fmt.Println("data does not exsist")
	}
}

//receiver function to get all customers record...
func (a Customer) getAll() dat {
	fmt.Println("all records.......")
	return data
}

//receiver function to get a particular customer data....
func (a Customer) get(k int) Customer {
	if _, ok := data[k]; ok {
		fmt.Println("displaying particular customer data ")
		return data[k]

	}
	return Customer{}
}

func main() {
	//multiple instance of Customer struct
	c1 := Customer{1, "sandhya", "bangalore", 9972072092}
	c2 := Customer{2, "shambavi", "chennai", 9972043768}
	c3 := Customer{3, "sangeetha", "hyderabad", 8197234789}
	fmt.Println(data)
	//intializing a map[] with created instances of Customer struct
	data[1] = c1
	data[2] = c2
	data[3] = c3
	fmt.Println(data)
	c.add(4, Customer{4, "sandhya", "bangalore", 6534218956})
	fmt.Println(data)
	c.update(1, Customer{1, "susmitha", "pune", 9972043768})
	//fmt.Println(data)
	// c.delete(7)
	fmt.Println(data)
	//c.delete(2)
	//fmt.Println(data)
	//fmt.Println(c.get(7))
	fmt.Println(c.getAll())

}
