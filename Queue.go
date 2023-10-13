package main

type Queue struct {
	data []int // массив элементов очереди
	size int  // максимальный размер очереди
	head int // индекс первого элемента в очереди (элемент, который будет удален при вызове dequeue)
	tail int // индекс следующего свободного места в очереди (элемент, который будет заполнен при вызове enqueue)
}

func NewQueue(size int) *Queue { // создание очереди заданного размера и типа данных 
	return &Queue{data: make([]int, size), size: size}
}

func (q *Queue) Enqueue(item int) { // добавление элемента в конец очереди 
	if q.tail == q.size { // если очередь заполнена 
		return 
	}
	
	q.data[q.tail] = item // добавление элемента 
	q.tail++ // увеличение индекса следующего свободного места 
}

func (q *Queue) Peek() int { // получение первого элемента очереди без его удаления 
	if q.head == q.tail { // если очередь пустая 
		return -1 
	}
	
	return q.data[q.head] // возвращаем первый элемент очереди 
}

func (q *Queue) Dequeue() int { // получение первого элемента очереди с его удалением 
	if q.head == q.tail { // если очередь пустая 
		return -1 
	}
	
	item := q.data[q.head] // получаем первый элемент очереди 
	q.data[q.head] = 0 // удаляем его из очереди 
	q.head++ // увеличиваем индекс первого элемента 
	
	return item // возвращаем удаленный элемент 
}

func (q *Queue) Len() int { // получение текущей длины очереди 
	return q.tail - q.head 
}