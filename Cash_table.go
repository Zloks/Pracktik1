package main

// Структура для элемента хэш-таблицы
type HashTableItem struct {
    key   string
    value int
}

// Структура для хэш-таблицы
type HashTable struct {
    items map[int]HashTableItem
}

// Функция для создания новой хэш-таблицы
func NewHashTable() *HashTable {
    table := HashTable{}
    table.items = make(map[int]HashTableItem)
    return &table
}

// Функция для получения хэша ключа
func (table *HashTable) hash(key string) int {
    sum := 0
    for _, char := range key {
        sum += int(char)
    }
    return sum % len(table.items)
}

// Функция для добавления элемента в хэш-таблицу
func (table *HashTable) Add(key string, value int) {
    item := HashTableItem{key, value}
    index := table.hash(key)
    table.items[index] = item
}

// Функция для получения значения по ключу из хэш-таблицы
func (table *HashTable) Get(key string) (int, bool) {
    index := table.hash(key)
    item, ok := table.items[index]
    if !ok {
        return 0, false
    }
    return item.value, true
}

// Функция для удаления элемента из хэш-таблицы
func (table *HashTable) Remove(key string) {
    index := table.hash(key)
    delete(table.items, index)
}