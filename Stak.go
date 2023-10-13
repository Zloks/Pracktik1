package main

type Stack struct {
	data []int
	size int // максимальный размер стека
	top  int // индекс верхнего элемента стека
}

func NewStack(size int) *Stack { // создание стека заданного размера и типа данных
	return &Stack{data: make([]int, size), size: size, top: -1}
}

func (s *Stack) Push(item int) { // добавление элемента на вершину стека
	if s.top == s.size-1 { // если стек заполнен
		return
	}

	s.top++              // увеличиваем индекс верхнего элемента
	s.data[s.top] = item // добавляем новый элемент на вершину стека
}

func (s *Stack) Peek() int { // получение верхнего элемента стека без его удаления
	if s.top == -1 { // если стек пустой
		return -1
	}

	return s.data[s.top] // возвращаем верхний элемент стека
}

func (s *Stack) Pop() int { // получение верхнего элемента стека с его удалением
	if s.top == -1 { // если стек пустой
		return -1
	}

	item := s.data[s.top] // получаем верхний элемент стека
	s.data[s.top] = 0     // удаляем его из стека
	s.top--               // уменьшаем индекс верхнего элемента

	return item // возвращаем удаленный элемент
}

func (s *Stack) Len() int { // получение текущей длины стека
	return s.top + 1
}
