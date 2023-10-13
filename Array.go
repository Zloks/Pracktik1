package main

type Array struct {
    data []int
}

func New_Array() *Array {
    return &Array{data: make([]int, 0)} // Создаём новый массив
}

func (a *Array) Add(item int) {
    a.data = append(a.data, item) // Добавляем элемент
}

func (a *Array) Get(index int) int { // Получаем элемент по индексу
    return a.data[index]
}

func (a *Array) Set(index int, item int) { // изменяем элемент по индексу
    a.data[index] = item
}

func (a *Array) Remove(index int) { // Удаление элемента
    a.data = append(a.data[:index], a.data[index+1:]...)
}

func (a *Array) Len() int { // Получаем текущую длинну массива
    return len(a.data)
}


