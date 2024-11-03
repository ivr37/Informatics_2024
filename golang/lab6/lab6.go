package lab6

<<<<<<< HEAD
import (
	"fmt"
)
=======
import "fmt"
>>>>>>> 0c6b35e (fixed lab4 and added lab6)

type Phone struct {
	Color  string
	OS     string
	Number string
}

func NewPhone(color, os, number string) *Phone {
	p := new(Phone)
	p.Color = color
	p.OS = os
	p.Number = number
	return p
}

func (p *Phone) SetNumber(number string) { p.Number = number }
<<<<<<< HEAD
func (p Phone) GetOS() string            { return p.OS }
func (p Phone) GetColor() string         { return p.Color }
func (p Phone) GetNumber() string        { return p.Number }

func RunLab6() {
	honor := NewPhone("Белый", "Андроид", "89203568472")
	honor.SetNumber(("89203657254"))
	fmt.Println("У нас есть телефон с операционной системой", honor.GetOS())
	fmt.Println("Цвет телефона:", honor.GetColor())
	fmt.Println("Номер телефона:", honor.GetNumber())
=======
func (p Phone) GetNumber() string        { return p.Number }
func (p Phone) GetOS() string            { return p.OS }
func (p Phone) GetColor() string         { return p.Color }

func RunLab6() {
	honor := NewPhone("Белый", "Андроид", "89014578541")
	honor.SetNumber("89203418794")
	fmt.Println("У нас есть мобильный телефон марки honor с операционной системой", honor.GetOS())
	fmt.Println("Цвет телефона:", honor.GetColor())
	fmt.Println("Номер телефона пользователя:", honor.GetNumber())
>>>>>>> 0c6b35e (fixed lab4 and added lab6)
}
