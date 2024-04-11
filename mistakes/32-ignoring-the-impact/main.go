package main

import "fmt"

// main #32: Ignoring the impact of using pointer elements in range loops
func main() {
	m := map[string]LargeStruct{
		"id": {foo: "foo"},
	}
	updateMapValue(m, "id")
	fmt.Println(m)

	m2 := map[string]*LargeStruct{
		"id": {foo: "foo"},
	}
	updateMapPointer(m2, "id")
	fmt.Printf("%+v\n", m2["id"])

	s := &Store{
		m: make(map[string]*Customer),
	}
	s.storeCustomers([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})
	fmt.Printf("%+v\n", s.m)
}

type LargeStruct struct {
	foo string
}

func updateMapValue(mapValue map[string]LargeStruct, id string) {
	value := mapValue[id]
	value.foo = "bar"
	// 如果不赋值，那么mapValue中的值不会被更新
	mapValue[id] = value
}
func updateMapPointer(mapPointer map[string]*LargeStruct, id string) {
	mapPointer[id].foo = "bar"
}

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

func (s *Store) storeCustomersWithFailed(customers []Customer) {
	for _, customer := range customers {
		// 这里都是指向同一个地址，所以最终map中的值都是相同的
		fmt.Printf("%p\n", &customer)
		s.m[customer.ID] = &customer
	}
}

func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		current := customer
		s.m[customer.ID] = &current
	}

	//for i := range customers {
	//	customer := &customers[i]
	//	s.m[customer.ID] = customer
	//}
}
