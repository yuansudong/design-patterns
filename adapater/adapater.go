package adapater

import "fmt"

type A struct {
}

func (a *A) DisplayA() {
	fmt.Println("ShowA")
}

type B struct {
}

func (a *B) DisplayB() {

	fmt.Println("ShowB")
}

type C struct {
}

// DisplayC
func (c *C) DisplayC() {
	fmt.Println("ShowC")
}

// Feture
type Feture struct {
	*C
	*B
	*A
}

type FT int32

const (
	F_A FT = iota + 1
	F_B
	F_C
)

func (f *Feture) Show(ty FT) {
	switch ty {
	case F_A:
		f.A.DisplayA()
	case F_B:
		f.B.DisplayB()
	case F_C:
		f.C.DisplayC()
	}
}
