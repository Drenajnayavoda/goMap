package main

import (
	"fmt"
	"map/map"
)

func main() {
	m := _map.NewMap()
	fmt.Println("для 12 ключ -", m.Add(12))
	m.Add(3)
	m.Add(66)
	m.Add(23)
	m.Print()
	fmt.Println("колличесво элементов в map -", m.Len())
	if value, ok := m.GetByIndex(3); ok {
		fmt.Println("значение по ключу '3' -", value)
	} else {
		fmt.Println("Элемент с ключом '3' не найден")
	}
	if id, ok := m.GetByValue(66); ok {
		fmt.Println("у значения '66' ключ -", id)
	} else {
		fmt.Println("Элемент со значением '66' не найден")
	}
	m.RemoveByValue(66)
	m.Print()
	m.RemoveByIndex(0)
	m.Print()
	m.Clear()
	m.Print()
}
