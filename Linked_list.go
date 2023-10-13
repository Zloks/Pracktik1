package main
import "fmt"
type LinkedList struct {
    data int
    next *LinkedList
}

func New_LinkedList(item int) *LinkedList {
    return &LinkedList{item, nil} // Создаём новый список
}

func (a *LinkedList) Add(item int) { // Добавляем элемент в список
    if(a.next == nil) { // Если если следующего элемента списка не существует 
      a.next = New_LinkedList(item) // Создаём новый элемент
      return  // Выходим из функции
    }
    (a.next).Add(item) // Рекурсивный вызов функции для следующего элемента списка
    
}

func (a *LinkedList) Get(index int) *LinkedList { // Получаем элемент по индексу
    if(a == nil) {return a} // Если список пустой возвращаем ничего 
    if (index == 0) { // Если индекс равен 0 возвращаем текущий элемент списка
        return a 
    }
    return (a.next).Get(index-1) // Рекурсивный вызов функции для следующего элемента списка
}

func (a *LinkedList) Set(index int, item int) bool { // изменяем элемент по индексу
    if(a == nil) {return false} // Если список пустой возвращаем ошибку 
    if (index == 0) { // Если индекс равен 0 изменяем текущий элемент списка
        a.data = item // Изменяем текущий элемент списка
        return true // Выходим из функции 
    }
    return (a.next).Set(index-1, item) // Рекурсивный вызов функции для следующего элемента списка
}

func (a *LinkedList) Remove(index int) bool { // Удаление элемента
  if(a == nil) {return false} // Если список пустой возвращаем ошибку
    if (index == 1 && a.next == nil) { // Если индекс ра��ен 1 и следующий элемент списка не существует возвращаем ошибку
      return false
    } else if (index == 1 && a.next != nil) { // Если индекс равен 1 и следующий элемент списка существует
      a.next = a.next.next  // Удаляем следующий элемент списка
      return true // Выходим из функции 
    }
    return (a.next).Remove(index-1) // Рекурсивный вызов функции для следующего элемента списка
}

func (a *LinkedList) Len() int { // Получаем текущую длинну массива
    if (a == nil) {return 0} // Если список пустой возвращаем 0
    if (a.next == nil){ return 1} // Если следующего элемента списка не существует возвращаем 1
    return 1 + a.next.Len() // Если следующего элемента списка существует возвращаем 1 + значение функции Len()
}

func (a *LinkedList) sprint(){ //Вывод всех элементов
  fmt.Print(a.data) // Вывод текущего элемента списка
  if (a.next != nil) {a.next.sprint()} // Рекурсивный вызов функции для следующего элемента списка
}