package practicaParcial2

//Resolver a mano el siguiente problema usando backtracking: Un Pastor tiene que atravesar a la otra orilla de un río con un Lobo, una Cabra y una Lechuga. Dispone de una barca en la que solo caben él y una de las otras tres cosas (es decir en cada viaje en barca puede ir el Pastor solo o el Pastor y alguno de los otros tres elementos). Si el Lobo se queda solo con la Cabra se la come, si la Cabra se queda sola con la Lechuga se la come. Pista: como estructuras de datos se puede usar un slice left, con todos los elementos que se encuentran en la orilla izquierda, un slice right, con todo lo que se encuentra a la derecha y barca que puede tomar los valores izq, o derecho, según donde se encuentre. Al comienzo left = [P, Lo, Le, C], right =[], barca = izq. Mostrar en cada paso como quedan las estructuras de datos. Por ejemplo si en el primer paso cruzan el Pastor y la Cabra las estructuras quedarían: left = [ Lo, Le], right =[P, C], barca = der.

/*
Se resuelve en foto.
*/