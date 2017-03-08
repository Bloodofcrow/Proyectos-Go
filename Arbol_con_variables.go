package main

import(
	io "fmt"
	str "strings"
)


//::::::::::::::::::: Nuevos Tipos de Datos ::::::::::::::

type Nodo struct{
	Valor int
	Nombre string
}

type Stack struct{
	nodos []*Nodo
	contador int
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::

//:::::::::::::::::::::: Fuciones ::::::::::::::::::::::::

//-----------------Información del Nodo------------------
func (nodo *Nodo) String() string{
	return nodo.Nombre
}
//-------------------------------------------------------

//------------------Funciones del Stack------------------
func NuevoStack() *Stack {
	return &Stack{}
}

func (pila *Stack) Push(nodo *Nodo){
	pila.nodos = append(pila.nodos[:pila.contador], nodo)
	pila.contador++
}

func (pila *Stack) Pop() *Nodo{
	if pila.contador == 0{
		return nil
	}
	pila.contador--
	return pila.nodos[pila.contador]
}
//-------------------------------------------------------

func ResolverPila(pila Pila) int{
	termino := pila.Pop()
	return 0
}

func main(){
	
	arbol1 := ":= + 5 3 x"
	arbol2 := ":= + - 5 2 4 y"
	arbol3 := ":= / x y z"
	pila1 := NuevoStack()
	pila2 := NuevoStack()
	pila3 := NuevoStack()
	pilaAux := NuevoStack()
	array1 := str.Split(arbol1, " ")
	array2 := str.Split(arbol2, " ")
	array3 := str.Split(arbol3, " ")
	for i:=0; i<len(array1); i++{
		pila1.Push(&Nodo{i,array1[i]})
	}
	for i:=0; i<len(array2); i++{
		pila2.Push(&Nodo{i,array2[i]})
	}
	for i:=0; i<len(array3); i++{
		pila3.Push(&Nodo{i,array3[i]})
	}

}
