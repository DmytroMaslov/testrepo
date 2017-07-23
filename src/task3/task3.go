package task3

import (
	"math"
	"sort"
	"errors"
)
type Triagle struct {
	Name string
	A float64
	B float64
	C float64
	area float64
}

func SortTriagle(triaglSlice []Triagle) []string{

	for _,t:=range triaglSlice {
		t.area = t.GetArea()
	}
	sort.Sort(sort.Reverse(byArea(triaglSlice)))
	name:=make([]string, len(triaglSlice))

	i := 0
	for _, v :=range triaglSlice{
		name[i] = v.Name
		i++
	}
	return name
}

func IsValid (triaglSlice []Triagle) error{
	for _,v := range triaglSlice{
		if !(v.isTriangle()) {
			return errors.New("It's not Triangle:"+ v.Name)
		}
		if (v.A<0|| v.B<0||v.C<0){
			return errors.New("Incorrect side:"+ v.Name)
		}
	}
	return nil
}

func (t *Triagle) GetArea () float64{
	p := (t.A+t.B+t.C)/2
	return math.Sqrt(p*(p-t.A)*(p-t.B)*(p-t.C))
}

func (t *Triagle) isTriangle() bool{
	if (t.A<t.B+t.C)&&(t.B<t.A+t.C)&&(t.C<t.B+t.A){
		return true
	} else {return false}
}

type byArea []Triagle

func(a byArea) Len() int {
	return len(a)
}
func (a byArea) Swap(i, j int){
	a[i], a[j] = a[j], a[i]
}
func (a byArea) Less(i, j int) bool {
	return a[i].area < a[j].area
	}