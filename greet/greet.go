package greet

// a nivel paquete no puede usarse con la abreviacion :=
// se comparte entre paquetes, todo lo que este declarado a nivel packete

//todo lo que este declarado con la primera letra en minuscula, hace que los demas
//paquetes no lo puedan usar (privado solo para su package)

//Si esta en mayusculas, puede usarse en cualquier otro packete

// es el equivalente a hacer un export
var emoji = "Pebete"



func English() string {
	return "hello" + emoji
}

func Italian() string {
	return "ciao" + emoji
}
