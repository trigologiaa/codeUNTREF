package practicaParcial1

//Agregar un método al TAD anterior que reciba como parámetro dos meses y devuelva los días que llovieron en ambos meses. Por ejemplo si en diciembre llovieron los días 1, 10, 15, 29, 30 y en enero los días 3, 5, 15, 29, 30 La respuesta es [15, 29, 30]

import(
	e "errors"
)

func (rl *RegistroLluvias) DiasLluviaEnAmbosMeses(mes1, mes2 string) ([]uint8, error) {
	bitMap1, existe1 := rl.meses[mes1]
	bitMap2, existe2 := rl.meses[mes2]
	if !existe1 || !existe2 {
		return nil, e.New("uno o ambos meses no son válidos")
	}
	var dias []uint8
	for indice := uint8(0); indice < 31; indice++ {
		on1, _ := bitMap1.IsOn(indice)
		on2, _ := bitMap2.IsOn(indice)
		if on1 && on2 {
			dias = append(dias, indice + 1)
		}
	}
	return dias, nil
}