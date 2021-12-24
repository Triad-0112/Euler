package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
}

const (
	mash          = 65
	alphalen      = 26
	numlimit      = 1000
	nameVaultSize = alphalen * alphalen * numlimit
)

var nameVault = generatenamepool(nameVaultSize)

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		err := r.Reset()
		if err != nil {
			return "", err
		}
	}
	return r.name, nil
}

func (r *Robot) Reset() error {
	if len(nameVault) == 0 {
		r.name = ""
		return errors.New("no more names available")
	}
	r.name = encryptTex(nameVault[0])
	nameVault = nameVault[1:]
	return nil
}
func generatenamepool(size int) []int {
	vault := make([]int, size)
	for i := 0; i < size; i++ {
		vault[i] = i
	}
	rand.Shuffle(size, func(i, j int) {
		vault[i], vault[j] = vault[j], vault[i]
	})
	return vault
}
func encryptTex(a int) string {
	alpha := a / numlimit
	mk1 := alpha / alphalen
	mk2 := alpha % alphalen
	mk3 := a % numlimit
	return fmt.Sprintf("%v%v%03d", encryptLet(mk1), encryptLet(mk2), mk3)
}
func encryptLet(a int) string {
	return string(rune(a + mash))
}
