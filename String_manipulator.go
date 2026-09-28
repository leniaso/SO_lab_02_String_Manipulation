package main

import (
	"fmt"
	"os"
)

func main() {
	entrada := getString()
	if entrada == "" {
		fmt.Println("No se recibió ninguna entrada válida.")
		return
	}

	entrada_rune := stringToRune(entrada)
	fmt.Println("Rune:", entrada_rune)
	entrada_rune = filtrarCaracteres(entrada_rune)
	copy_runa := copiarRunas(entrada_rune)
	numVocales, numA, numE, numI, numO, numU, numConsonantes := contarVocalesYConsonantes(entrada_rune)
	fmt.Println("Vocales:", numVocales, "a:", numA, "e:", numE, "i:", numI, "o:", numO, "u:", numU, "Consonantes:", numConsonantes)
	fmt.Println(string(entrada_rune))
	fmt.Println(string(revertirCaracteres(entrada_rune)))
	fmt.Println(string(cambiarEspacios(copy_runa)))
}

func getString() string {
	if len(os.Args) < 2 {
		fmt.Println("Error: no se proporcionó ningún argumento.")
		return ""
	}
	entrada := os.Args[1]
	if len(entrada) > 100 {
		fmt.Println("Error: La entrada supera los 100 caracteres.")
		return ""
	}
	return entrada
}

func stringToRune(s string) []rune {
	return []rune(s)
}

func filtrarCaracteres(runas []rune) []rune {
	idxEscritura := 0

	for idxLectura := 0; idxLectura < len(runas); idxLectura++ {
		ptrLectura := &runas[idxLectura]

		esLetra := (*ptrLectura >= 'A' && *ptrLectura <= 'Z') || (*ptrLectura >= 'a' && *ptrLectura <= 'z')
		esConTilde := (*ptrLectura == 'á' || *ptrLectura == 'é' || *ptrLectura == 'í' || *ptrLectura == 'ó' || *ptrLectura == 'ú' || *ptrLectura == 'Á' || *ptrLectura == 'É' || *ptrLectura == 'Í' || *ptrLectura == 'Ó' || *ptrLectura == 'Ú')
		esEspacio := (*ptrLectura == ' ')

		if esLetra || esConTilde || esEspacio {
			ptrEscritura := &runas[idxEscritura]
			*ptrEscritura = *ptrLectura
			idxEscritura++
		}

	}
	return runas[:idxEscritura]
}

func revertirCaracteres(runas []rune) []rune {
	idxInicio := 0
	idxFin := len(runas) - 1

	for idxInicio < idxFin {
		ptrInicio := &runas[idxInicio]
		ptrFin := &runas[idxFin]
		*ptrInicio, *ptrFin = *ptrFin, *ptrInicio
		idxInicio++
		idxFin--
	}
	return runas
}

func cambiarEspacios(runas []rune) []rune {
	idxEscritura := 0
	for idxLectura := 0; idxLectura < len(runas); idxLectura++ {

		ptrLectura := &runas[idxLectura]
		esEspacio := (*ptrLectura == ' ')
		if esEspacio {
			idxEscritura = idxLectura
			ptrEscritura := &runas[idxEscritura]
			*ptrEscritura = '_'
		}

	}
	return runas
}

func copiarRunas(original []rune) []rune {
	copia := make([]rune, len(original))
	copy(copia, original)
	return copia
}

func contarVocalesYConsonantes(runas []rune) (numVocales, numA, numE, numI, numO, numU, numConsonantes int) {
	for idxLectura := 0; idxLectura < len(runas); idxLectura++ {
		ptrLectura := &runas[idxLectura]
		c := *ptrLectura

		if c == ' ' {
			// los espacios ya no cuentan como letra, se saltan
			continue
		}

		switch c {
		case 'a', 'A', 'á', 'Á':
			numA++
		case 'e', 'E', 'é', 'É':
			numE++
		case 'i', 'I', 'í', 'Í':
			numI++
		case 'o', 'O', 'ó', 'Ó':
			numO++
		case 'u', 'U', 'ú', 'Ú':
			numU++
		default:
			// como el slice ya pasó por filtrarCaracteres, lo único que
			// puede caer aquí es una consonante
			numConsonantes++
		}
	}

	numVocales = numA + numE + numI + numO + numU
	return
}
