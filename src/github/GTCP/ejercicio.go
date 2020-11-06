package main

import "fmt" 
import "errors"

type OrganizacionProtectoraDeAnimales struct {
	nombre  string
	empleados []persona
	animales []animal
}
type animal struct {
	chip int
    name string
	tipo string
	raza string
    edad int
	mesEntrada int
	añoEntrada int
}
type persona struct {
	documento int
	name string
	cargo string
    edad int
	salario int
	horas int
}
func (O *OrganizacionProtectoraDeAnimales) agregarEmpleado(documento int,name string,cargo string, edad int, salario int,horas int) {
	incorporacion := persona{documento,name,cargo,edad,salario,horas}
	O.empleados = append(O.empleados, incorporacion)
}
func (O *OrganizacionProtectoraDeAnimales) agregarAnimal(id int,name string, tipo string,raza string, edad int, mesEntrada int,añoEntrada int) {
	incorporacion := animal{id, name, tipo,raza, edad,mesEntrada,añoEntrada}
	O.animales = append(O.animales, incorporacion)
}
func (O *OrganizacionProtectoraDeAnimales) obtenerAnimal(chip int) (animal, error) {
	var animal animal
	for _, ani := range O.animales {
		if ani.chip == chip {
			animal = ani
			return animal, nil
		}
	}
	return animal, errors.New("No se encontró animal")
}
func (O *OrganizacionProtectoraDeAnimales) obtenerEmpleado(documento int) (persona, error) {
	var persona persona
	for _, pers := range O.empleados {
		if pers.documento == documento {
			persona = pers
			return persona, nil
		}
	}
	return persona, errors.New("No se encontró empleado")
}
func (E *OrganizacionProtectoraDeAnimales) obtenerAnimales() []animal {
	return E.animales
}
func (E *OrganizacionProtectoraDeAnimales) obtenerEmpleados() []persona {
	return E.empleados
}
func (O *OrganizacionProtectoraDeAnimales) adoptarAnimal(chip int) {
	for i, animal := range O.animales {
		if animal.chip == chip {
			O.animales = append(O.animales[:i], O.animales[i+1:]...)
		}
	}
}
func (O *OrganizacionProtectoraDeAnimales) despedirEmpleado(documento int) {
	for i, persona := range O.empleados {
		if persona.documento == documento {
			O.empleados = append(O.empleados[:i], O.empleados[i+1:]...)
		}
	}
}
func (O *OrganizacionProtectoraDeAnimales) actualizarEmpleado(documento int,name string,cargo string, edad int, salario int,horas int) {
	for i := 0; i < len(O.empleados); i++ {
		if O.empleados[i].documento == documento {
			O.empleados[i].name = name
			O.empleados[i].cargo = cargo
			O.empleados[i].edad = edad
			O.empleados[i].salario = salario
			O.empleados[i].horas = horas
		}
	}
}
func (O *OrganizacionProtectoraDeAnimales) actualizarAnimal(chip int,name string, tipo string,raza string, edad int, mesEntrada int,añoEntrada int) {
	for i := 0; i < len(O.animales); i++ {
		if O.animales[i].chip == chip {
			O.animales[i].name = name
			O.animales[i].tipo = tipo
			O.animales[i].raza = raza
			O.animales[i].edad = edad
			O.animales[i].mesEntrada = mesEntrada
			O.animales[i].añoEntrada = añoEntrada
		}
	}
}
func main(){
	var empleados = []persona{
		{	
			documento: 123456789,
			name: "alexia",
			cargo: "encargada",
			edad: 31,
			salario: 20000,
			horas: 4,
		},
		{
			documento: 987654321,
			name: "carlos",
			cargo: "empleado",
			edad: 23,
			salario: 30000,
			horas: 8,
		},
		{
			documento: 456789347,
			name: "amanda",
			cargo: "empleada",
			edad: 23,
			salario: 15000,
			horas: 4,
		},
    }
    var angeles = []animal{
		{	
			chip: 1,
			name: "blake",
			tipo: "perro",
			raza: "bulldog frances",
			edad: 8,
			mesEntrada: 3,
			añoEntrada: 2019,
		},
		{
			chip: 2,
			name: "milu",
			tipo: "gato",
			raza: "toyger",
			edad: 5,
			mesEntrada: 4,
			añoEntrada: 2020,	
		},
		{
			chip: 3,
			name: "hector",
			tipo: "raton",
			raza: "Pyromys",
			edad: 4,
			mesEntrada: 5,
			añoEntrada: 2019,	
		},
		{
			chip: 4,
			name: "melanie",
			tipo: "loro",
			raza: "Loris Arcoiris",
			edad: 25,
			mesEntrada: 9,
			añoEntrada: 2018,
		},
    }
	OrganizacionProtectoraDeAnimales := OrganizacionProtectoraDeAnimales{"Salvando Angeles peludos",empleados, angeles}

	OrganizacionProtectoraDeAnimales.agregarEmpleado(654897123,"aldo","empleado", 18, 20000, 4)

	OrganizacionProtectoraDeAnimales.agregarAnimal(5,"oscar", "cerdo","Large white", 2, 7, 2019)
	OrganizacionProtectoraDeAnimales.agregarAnimal(6,"hernesto", "caballo","American Morgan", 2, 7, 2019)
	OrganizacionProtectoraDeAnimales.agregarAnimal(7,"arnold", "perro","bulldog ingles", 2, 7, 2019)
	OrganizacionProtectoraDeAnimales.agregarAnimal(8,"manu", "perro","salchicha", 3, 4, 2020)

	//pido animal por chip rastreador
	fmt.Println("Traigo un animal por chip rastreador")
	animal1, error := OrganizacionProtectoraDeAnimales.obtenerAnimal(1)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(animal1)
	}
	fmt.Println("------------------")
	fmt.Println("Traigo todos los animales")

	animales := OrganizacionProtectoraDeAnimales.obtenerAnimales();
	for _, a := range animales {
		fmt.Println(a)
	}

	fmt.Println("------------------")
	//modifico atributos de empleado
	fmt.Println("Modifico atributos de un empleado")

	OrganizacionProtectoraDeAnimales.actualizarEmpleado(987654321, "carlos", "administrador", 18,30000,8)
	empleado1, error := OrganizacionProtectoraDeAnimales.obtenerEmpleado(987654321)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(empleado1)
	}
	fmt.Println("------------------")
	//despido a un empleado
	fmt.Println("despido un empleado por robar comida para perros")

	OrganizacionProtectoraDeAnimales.despedirEmpleado(654897123)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(empleado1)
	}
	fmt.Println("------------------")
	fmt.Println("traigo todos los empleados")

	empleados2 := OrganizacionProtectoraDeAnimales.obtenerEmpleados();
	for _, e := range empleados2 {
		fmt.Println(e)
	}
	fmt.Println("------------------")
}