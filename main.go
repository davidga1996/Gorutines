package main

import (
	"fmt"
	"time"
)

var stop int64

type Process struct {
	Id int64
	I  int64
}

type Processes struct{ SliceProcesses []Process }

func (ps *Processes) add(p Process) {
	ps.SliceProcesses = append(ps.SliceProcesses, p)

	for iterator := 0; iterator < len(ps.SliceProcesses); iterator = iterator + 1 {
		go ps.SliceProcesses[iterator].do()
	}
}

func (ps *Processes) shows() {
	for i := 0; i < len(ps.SliceProcesses); i = i + 1 {
		go ps.SliceProcesses[i].show()
	}
}

func (ps *Processes) remove(b int64) {
	var posicion int

	posicion = -1

	for i := 0; i < len(ps.SliceProcesses) && posicion == -1; i = i + 1 {
		if ps.SliceProcesses[i].Id == b {
			posicion = i
		}
	}

	if posicion != -1 {
		if posicion == len(ps.SliceProcesses)-1 {
			ps.SliceProcesses = append(ps.SliceProcesses[:posicion])
		} else {
			ps.SliceProcesses = append(ps.SliceProcesses[:posicion], ps.SliceProcesses[posicion+1:]...)
		}
	}
}

func (p *Process) do() {
	for {
		p.I = p.I + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Process) show() {
	for {
		fmt.Println("id ", p.Id, ": ", p.I)
		time.Sleep(time.Millisecond * 500)

		if stop == 1 {
			return
		}
	}
}

func main() {
	var option int64
	var id int64
	var removeItem int64
	id = 0

	Processes := Processes{}

	for true {

		fmt.Println("Menu: ")
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar Procesos")
		fmt.Println("3) Terminar proceso")
		fmt.Println("4) Salir")
		fmt.Print("Opcion: ")
		fmt.Scanln(&option)
		fmt.Println("")

		if option == 1 {
			Processes.add(Process{Id: id, I: 0})
			id = id + 1

		} else if option == 2 {
			//stop = stop == 0 ? 1 : 0
			if stop == 0 {
				stop = 1
			} else {
				stop = 0
			}
			Processes.shows()
		} else if option == 3 {
			fmt.Scanln(&removeItem)
			Processes.remove(removeItem)
		} else if option == 4 {
			return
		} else {
			fmt.Println("Opcion incorrecta")
		}
		fmt.Println("")
	}
}
