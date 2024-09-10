package parcial1MAMI

//Dada la anterior implementación de una lista enlazada simple, pero que debe mantener el orden de sus elementos. Se pide. Implementar el método Insertar, que reciba como parámetro un elemento y lo inserte en la posición que corresponda. Justificar el orden de Insertar

type node struct {
	value int
	next *node
}

func newNode(value int) *node {
	return &node {
		value: value,
		next: nil,
	}
}

type OrdLinkList struct {
	head *node
	tail *node
	size int
}

func NewOrdLinkList() *OrdLinkList {
	return &OrdLinkList {
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *OrdLinkList) Insertar(value int) {
    nuevoNodo := newNode(value)
    l.size++
    if l.head == nil {
        l.head = nuevoNodo
        l.tail = nuevoNodo
        return
    }
    if value < l.head.value {
        nuevoNodo.next = l.head
        l.head = nuevoNodo
        return
    }
    if value >= l.tail.value {
        l.tail.next = nuevoNodo
        l.tail = nuevoNodo
        return
    }
    actual := l.head
    for actual.next != nil && value >= actual.next.value {
        actual = actual.next
    }
    nuevoNodo.next = actual.next
    actual.next = nuevoNodo
}
/*
El método Insert está diseñado para mantener el orden ascendente de los elementos en la lista enlazada. Para lograr esto:

Caso 1: Si la lista está vacía, el nuevo nodo se convierte tanto en head como en tail.

Caso 2: Si el valor a insertar es menor que el valor del nodo en head, el nuevo nodo se convierte en el nuevo head, manteniendo así el orden.

Caso 3: Si el valor a insertar es mayor o igual que el valor del nodo en tail, el nuevo nodo se convierte en el nuevo tail, asegurando que los elementos más grandes queden al final de la lista.

Caso 4: Si el valor a insertar está entre dos nodos existentes, recorremos la lista para encontrar la posición adecuada donde el nuevo nodo debe ser insertado, manteniendo el orden ascendente.

En todos los casos, el método Insert garantiza que los elementos se inserten en la posición correcta para mantener el orden ascendente de los valores en la lista enlazada simple OrdLinkList[int].

Esta implementación asegura eficiencia en la inserción, siendo O(1) en el mejor caso (inserción al principio o final) y O(n) en el peor caso (inserción en el medio, donde n es el número de elementos en la lista hasta el punto de inserción).
*/