/*Autores:
Jehison Andres Rodriguez
Juan Pablo Soler
Cesar Mauricio Forero
*/

package main

import (
	bu "bufio"
	io "fmt"
	os "os"
	conv "strconv"
	str "strings"
)

//::::::::::::::::::: Nuevos Tipos de Datos ::::::::::::::

type Nodo struct {
	Valor  int
	Nombre string
}

type Stack struct {
	nodos    []*Nodo
	contador int
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::

//:::::::::::::::::::::: Fuciones ::::::::::::::::::::::::

//-----------------Informacion del Nodo------------------
func (nodo *Nodo) String() string {
	return nodo.Nombre
}

//-------------------------------------------------------

//------------------Funciones del Stack------------------
func NuevoStack() *Stack {
	return &Stack{}
}

func (pila *Stack) Push(nodo *Nodo) {
	pila.nodos = append(pila.nodos[:pila.contador], nodo)
	pila.contador++
}

func (pila *Stack) Pop() *Nodo {
	if pila.contador == 0 {
		return nil
	}
	pila.contador--
	return pila.nodos[pila.contador]
}

//-------------------------------------------------------

func ResolverPila(pila *Stack) int {
	pilaAux := NuevoStack()
	rsta := 0
	for i := 0; i < len(pila.nodos); i++ {
		termino := pila.Pop()
		aux, err := conv.Atoi(termino.Nombre)
		if err != nil {
			switch termino.Nombre {
			case "+":
				rsta = pilaAux.Pop().Valor + pilaAux.Pop().Valor
				pilaAux.Push(&Nodo{rsta, ""})
			case "-":
				rsta = pilaAux.Pop().Valor - pilaAux.Pop().Valor
				pilaAux.Push(&Nodo{rsta, ""})
			case "/":
				numerador := pilaAux.Pop().Valor
				denominador := pilaAux.Pop().Valor
				if denominador != 0 {
					rsta = numerador / denominador
				} else {
					rsta = numerador / 1
				}
				pilaAux.Push(&Nodo{rsta, ""})
			case "*":
				rsta = pilaAux.Pop().Valor * pilaAux.Pop().Valor
				pilaAux.Push(&Nodo{rsta, ""})
			}
		} else {
			pilaAux.Push(&Nodo{aux, ""})
		}
	}
	return rsta
}

func LecturaDesdeConosola() (string, error) {
	lectura := bu.NewReader(os.Stdin)
	s, err := lectura.ReadString('\n')

	return str.TrimSpace(s), err
}

func main() {

	io.Print("Ingrese la expresion (Arbol 1) en postfijo: ")
	arbol1String, err := LecturaDesdeConosola()

	io.Print("Ingrese la expresion (Arbol 2) en postfijo: ")
	arbol2String, err := LecturaDesdeConosola()

	io.Print("Ingrese la expresion (Arbol 3) en postfijo: ")
	arbol3String, err := LecturaDesdeConosola()

	pila1 := NuevoStack()
	pila2 := NuevoStack()
	pila3 := NuevoStack()

	array1 := str.Split(arbol1String, " ")
	array2 := str.Split(arbol2String, " ")
	array3 := str.Split(arbol3String, " ")

	for i := 0; i < len(array1); i++ {
		pila1.Push(&Nodo{i, array1[i]})
	}
	for i := 0; i < len(array2); i++ {
		pila2.Push(&Nodo{i, array2[i]})
	}
	for i := 0; i < len(array3); i++ {
		pila3.Push(&Nodo{i, array3[i]})
	}

	x := ResolverPila(pila1)
	y := ResolverPila(pila2)
	arr := [2]int{x, y}
	for i := 0; i <= len(arr); i++ {
		pila3.Pop()
	}
	for i := 0; i < len(arr); i++ {
		pila3.Push(&Nodo{i, conv.Itoa(arr[i])})
	}
	z := ResolverPila(pila3)
	io.Println("X=", arr[0], " Y=", arr[1], " Z=", z)

	if err != nil {
		io.Println("Error cuando se escaneo la entrada", err.Error())
		return
	}

}
