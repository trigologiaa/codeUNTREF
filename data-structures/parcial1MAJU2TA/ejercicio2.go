package parcial1MAJU2TA

//Implementar una primitiva (método) de BitMap, llamada String, que devuelva una cadena de caracteres con la representación binaria del mapa del bit (la cadena siempre debe tener 32 caracteres, ceros y/o unos).

import (
	"errors"
	"strings"
)

const (
	BitmapSize uint8 = 32
)

type BitMap struct {
	bits uint32
}

func NewBitMap() *BitMap {
	return &BitMap{bits: 0b0}
}

func (bm *BitMap) On(pos uint8) error {
	if isOutOfRange(pos) {
		return errors.New("posición no válida")
	}

	bm.bits |= 0b1 << pos

	return nil
}

func (bm *BitMap) Off(pos uint8) error {
	if isOutOfRange(pos) {
		return errors.New("posición no válida")
	}

	bm.bits &= ^(0b1 << pos)

	return nil
}

func (bm *BitMap) IsOn(pos uint8) (bool, error) {
	if isOutOfRange(pos) {
		return false, errors.New("posición no válida")
	}

	return bm.bits&(1<<pos) != 0b0, nil
}

func (bm *BitMap) GetMap() uint32 {
	return bm.bits
}

func isOutOfRange(pos uint8) bool {
	return pos >= BitmapSize
}

func(bm *BitMap) String() string {
	var resultado strings.Builder
	for i := BitmapSize - 1; i > 0; i-- {
		mask := uint32(1 << i)
		if bm.bits&mask != 0 {
			resultado.WriteByte('1')
		} else {
			resultado.WriteByte('0')
		}
	}
	return resultado.String()
}