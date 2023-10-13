package main
import "fmt"
type LinkedList2 struct {
    data int
    next *LinkedList2
    prev *LinkedList2    
}

func New_LinkedList2(item int) *LinkedList2 {
    return &LinkedList2{item, nil, nil}  // Создаём новый список
}

func (a *LinkedList2) Add2(item int) { // Добавляем элемент в список
    if(a.next == nil) { // Если список пустой
      a.next = New_LinkedList2(item) // Создаём новый элемент
      a.next.prev = a // Связываем новый элемент с предыдущим
      return  // Выходим из функции
    }
    (a.next).Add2(item) // Добавляем новый элемент в список
    
}

func (a *LinkedList2) Get2(index int) *LinkedList2 { // Получаем элемент по индексу
    if(a == nil) {return a} // Если список пустой возвращаем ничего 
    if (index == 0) { // Если индекс равен 0 возвращаем первый элемент списка
        return a
    }
    return (a.next).Get2(index-1) 
}

func (a *LinkedList2) Set2(index int, item int) bool { // изменяем элемент по индексу
    if(a == nil) {return false} // Если список пустой возвращаем ошибку
    if (index == 0) { // Если индекс равен 0 изменяем первый элемент списка
        a.data = item // Изменяем первый элемент списка
        return true // Выходим из функции
    }
    return (a.next).Set2(index-1, item) // Изменяем элементы в списке
}

func (a *LinkedList2) Remove2(index int) bool { // Удаление элемента
  if(a == nil) {return false} // Если список пустой возвращаем ошибку
    if (index == 1 && a.next == nil) { // Если индекс равен 1 и следующий элемент списка пусто возвращаем ошибку
      return false
    } else if (index == 1 && a.next != nil) { // Если индекс равен 1 и следующий элемент списка не пуст
      a.next = a.next.next // Удаляем следующий элемент списка
      return true 
    }
    return (a.next).Remove2(index-1) // 
}

func (a *LinkedList2) Len2() int { // Получаем текущую длинну массива
    //if (a == nil) return 0
    if (a.next == nil){ return 1}
    return 1 + a.next.Len2()
}

func (a *LinkedList2) sprint2() { //Вывод всех элементов
  fmt.Println(a.data) // Вывод текущего элемента списка
  if (a.next != nil) {a.next.sprint2()} // Рекурсивный вызов функции для следующего элемента списка
}
func (a *LinkedList2) Ad_b_index(index int, item int) bool { // Добавляем элемент в список по индексу
    if(a == nil) {return false} // Если список пустой или число элементов меньше индекса возвращаем ошибку
    if (index == 0) { // Если индекс равен 0 
        var b = New_LinkedList2(item)// Создаём новый элемент
        b.prev = a // Связываем новый элемент с предыдущим
        b.next = a.next // Связываем новый элемент с следующим
        a.next.prev = b // Связываем следующий элемент с новым
        a.prev.next = b // Связываем предыдущий элемент с новым
        return true
    }
    return (a.next).Ad_b_index(index-1, item) // Рекурсивный вызов функции для следующего элемента списка

}