package parcial1MAJU2MA

//Escribir una función que reciba como parámetro un diccionario cuyas claves son títulos de libros (strings) y cada clave tiene asociada un conjunto de autores (string). La función debe devolver un diccionario cuyas claves sean los autores y cuyo valor una lista enlazada simple de títulos de libros.

import(
	d "github.com/trigologiaa/data-structures/dictionary"
	l "github.com/trigologiaa/data-structures/list"
)

func AutoresYTitulos(libroDiccionario *d.Dictionary[string, []string]) *d.Dictionary[string, *l.LinkedList[string]] {
    autorDiccionario := d.NewDictionary[string, *l.LinkedList[string]]()
    for _, clave := range libroDiccionario.Keys() {
        autores, _ := libroDiccionario.Get(clave)
        for _, autor := range autores {
            if autorDiccionario.Contains(autor) {
                ListaDeAutores, _ := autorDiccionario.Get(autor)
                ListaDeAutores.Append(clave)
            } else {
                nuevaLista := l.NewLinkedList[string]()
                nuevaLista.Append(clave)
                autorDiccionario.Put(autor, nuevaLista)
            }
        }
    }
	return autorDiccionario
}