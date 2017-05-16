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
	reg "regexp"
)

var(
	lista Lista
) 

//::::::::::::::::::: Nuevos Tipos de Datos ::::::::::::::

type NodoLista struct {
	Valor  string
	Nombre string
}

type Lista struct {
	nodos    []*NodoLista
	contador int
}

type Nodo struct {
	Valor  int
	Nombre string
}

type Stack struct {c
	nodos    []*Nodo
	contador int
}

type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
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

//------------------Funciones de Lista------------------
func NuevaLista() *Lista {
	return &Lista{}
}

func (pila *Lista) PushLista(nodo *NodoLista) {
	pila.nodos = append(pila.nodos[:pila.contador], nodo)
	pila.contador++
}

func (pila *Lista) PopLista() *NodoLista {
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
				denominador := pilaAux.Pop().Valor
				if denominador != 0 {
					rsta = pilaAux.Pop().Valor / denominador
				} else {
					rsta = pilaAux.Pop().Valor / 1
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

//Función para la lectira de consola.
func LecturaDesdeConosola() (string, error) {
	lectura := bu.NewReader(os.Stdin)
	s, err := lectura.ReadString('\n')

	return str.TrimSpace(s), err
}

//Función para comprobar los strings.
func Comprobar(cadena string) bool {

	reg1, _ := reg.Compile(`\D`) //Valida que tenga algo, aparte de números.
	reg2, _ := reg.Compile(`\W`) //Valida que no tenga carácteres especiales.
	reg3, _ := reg.Compile(`\w`) //Valida que no tenga números o letras.

	if(reg1.MatchString(cadena) == true){

		if (reg2.MatchString(cadena) == false){

			//Sería un token "Variable".
			//Método para almacenar el token en la lista.
			lista.PushLista(&NodoLista{cadena, "Variable"})
			return true

		}else {

			if (reg3.MatchString(cadena) == false){

				//Sólo tiene carácteres especiales.
				//Método para almacenar el token en la lista.
				lista.PushLista(&NodoLista{cadena, "Operador"})
				return true

			}

		}

	}else if (reg3.MatchString(cadena) == false){

		//Método para almacenar el token en la lista.
		//Es una constante.
		lista.PushLista(&NodoLista{cadena, "Constante"})
		return true

	}

	return false

}

//Función que creará el árbol a partir de la lista de tokens.
func CrearArbol(lista *Lista) *Arbol {

//Método para crear el árbol de expresiones.

}

//Función para inicializar la pila que será la lista de tokens.
func InicializarPila(){

	lista := NuevaLista	()

}

func main() {

	InicializarPila()
	comprobado := true

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

	//Falta comprobar
	for i := 0; i < len(array1); i++ {
		if Comprobar(array1[i]) == true{
			pila1.Push(&Nodo{i, array1[i]})
		} else{
			comprobado = false
		}
	}

	lista1 := lista
	InicializarPila()
	arbol1 := CrearArbol(lista1)

	for i := 0; i < len(array2); i++ {
		if Comprobar(array2[i]) == true{
			pila1.Push(&Nodo{i, array2[i]})
		} else{
			comprobado = false
		}
	}

	lista2 := lista
	InicializarPila()
	CrearArbol(lista2)

	for i := 0; i < len(array3); i++ {
		if Comprobar(array3[i]) == true{
			pila1.Push(&Nodo{i, array3[i]})
		} else{
			comprobado = false
		}
	}	

	lista3 := lista
	InicializarPila()
	CrearArbol(lista3)

	//-----------------

	if comprobado == false{
		return
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
