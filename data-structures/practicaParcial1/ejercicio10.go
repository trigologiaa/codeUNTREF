package practicaParcial1

//Implementar un registro de lluvias anuales utilizando mapas de bits. El registro debe contar con un método que reciba un mes y la lista de los días que llovieron y registrarlo. Por ejemplo Registrar ("enero", 12, 22, 23, 28). También debe contar con un método que reciba como parámetro un mes y devuelva los días que llovieron en ese mes, por ejemplo Lluvias("enero") debe devolver [12, 22, 23, 28].

import(
	e "errors"
	b "github.com/trigologiaa/data-structures/bitmap"
)

type RegistroLluvias struct {
	meses map[string]	*b.BitMap
}

func NewRegistroLluvias() *RegistroLluvias {
	return &RegistroLluvias {
		meses: map[string]*b.BitMap {
			"enero":     	b.NewBitMap(),
			"febrero":   	b.NewBitMap(),
			"marzo":     	b.NewBitMap(),
			"abril":     	b.NewBitMap(),
			"mayo":      	b.NewBitMap(),
			"junio":     	b.NewBitMap(),
			"julio":     	b.NewBitMap(),
			"agosto":    	b.NewBitMap(),
			"septiembre":	b.NewBitMap(),
			"octubre":   	b.NewBitMap(),
			"noviembre": 	b.NewBitMap(),
			"diciembre": 	b.NewBitMap(),
		},
	}
}

func (rl *RegistroLluvias) Registrar(mes string, dias ...uint8) error {
	bitMap, existe := rl.meses[mes]
	if !existe {
		return e.New("mes no válido")
	}
	for _, dia := range dias {
		if dia < 1 || dia > 31 {
			return e.New("día no válido")
		}
		bitMap.On(dia - 1)
	}
	return nil
}

func (rl *RegistroLluvias) Lluvias(mes string) ([]uint8, error) {
	bitMap, existe := rl.meses[mes]
	if !existe {
		return nil, e.New("mes no válido")
	}
	var dias []uint8
	for indice := uint8(0); indice < 31; indice++ {
		on, _ := bitMap.IsOn(indice)
		if on {
			dias = append(dias, indice + 1)
		}
	}
	return dias, nil
}