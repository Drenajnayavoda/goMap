package _map

import "fmt"

type Map struct {
	HashTable map[int64]int64
	len       int64
}

// NewMap создает новую пустую карту
func NewMap() *Map {
	hashTable := make(map[int64]int64, 0)
	return &Map{HashTable: hashTable}
}

// Len возвращает количество элементов в карте
func (m *Map) Len() int64 {
	return m.len
}

// Add добавляет элемент в карту и возвращает его индекс (ключ)
func (m *Map) Add(value int64) int64 {
	key := m.len // Используем текущую длину карты как новый ключ
	m.HashTable[key] = value
	m.len++
	return key
}

// RemoveByIndex удаляет элемент из карты по ключу (индексу)
func (m *Map) RemoveByIndex(id int64) bool {
	if id < 0 || id >= m.len {
		fmt.Println("Invalid index")
		return false
	}

	if id == m.len-1 {
		delete(m.HashTable, id)
		m.len--
		return true
	}

	for i := id + 1; i < m.len; i++ {
		m.HashTable[i-1] = m.HashTable[i]
	}

	delete(m.HashTable, m.len-1)
	m.len--
	return true
}

// RemoveByValue удаляет элемент из карты по значению
func (m *Map) RemoveByValue(value int64) bool {

	for i := int64(0); i < m.len; i++ {
		if m.HashTable[i] == value {
			m.RemoveByIndex(i)
			return true
		}
	}
	return false
}

// GetByIndex возвращает значение элемента по ключу (индексу).
//
// Если элемента с таким ключом нет, то возвращается 0 и false
func (m *Map) GetByIndex(key int64) (int64, bool) {
	value, ok := m.HashTable[key]
	return value, ok
}

// GetByValue возвращает ключ (индекс) первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (m *Map) GetByValue(value int64) (int64, bool) {
	for key, val := range m.HashTable {
		if val == value {
			return key, true
		}
	}
	return 0, false
}

// Clear очищает карту
func (m *Map) Clear() {
	m.HashTable = make(map[int64]int64)
	m.len = 0
}

// Print выводит все элементы карты в консоль
func (m *Map) Print() {
	var i int64 = 0
	for ; i < m.len-1; i++ {
		fmt.Print(m.HashTable[i], "[", i, "], ")
	}
	fmt.Print(m.HashTable[i], "[", i, "]")
	fmt.Println("")
}
