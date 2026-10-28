package contador

import (
	"errors"
	"fmt"
)

func Canasta(puntos int) (string, error) {
	if puntos!=1 && puntos!=2 && puntos!=3{
		return "",errors.New("Los puntos deben ser 1 2 o 3")
	} 
	mensaje := fmt.Sprintf("Canasta de %d puntos\n",puntos)
	return mensaje,nil
}

func Contar(canasta []int)(int, error){
	total := 0 
	for _,puntos := range canasta{
		_,err := Canasta(puntos)
		if err != nil{
			return 0, err
		}
		total = total + puntos
	}
	return total,nil
}