//Task1
// {"Lenght": 10, "Width": 10, "Symbol": "S"}
package task1

import (
	"errors"
	"fmt"
	"encoding/json"
)

type ChessBoard struct {
	Width int `json:"width"`
	Height  int `json:"height"`
	Symbol string `json:"symbol"`
}
const SPACE = " "

func Run(param []byte) (str string, err error){
	var ch = new (ChessBoard)
	err = json.Unmarshal(param, &ch)
	if err != nil{
		return str, errors.New("can't unmarshal data")
	}
	if err = IsValid(ch.Height, ch.Width, ch.Symbol); err != nil{
		return str, err
	}
	return writeGrid(ch.Height, ch.Width, ch.Symbol), nil
}
func IsValid (l int, w int, c string) (error){
	if l< 2 {
		return errors.New("incorrect Width")
	}
	if w <2{
		return errors.New("incorrect Height")
	}
	if len(c) == 0{
		return errors.New("empty char symbol")
	}
	return nil
}

func writeGrid (l int, w int, c string) string{
	i:= 0
	var space = SPACE
	var char = c
	var newLine = "\n"
	var grid = make([]string, 0)
	for i<l {
		j :=0
		for j<w {
			if i%2 == 0{
				if j%2 == 0{
					grid = append(grid, char)
				} else {
					grid = append(grid, space)
				}

			}
			if i%2 == 1 {
				if j%2 == 0{
					grid = append(grid, space)
				} else {
					grid = append(grid, char)
				}
			}
			j++
		}
		grid = append(grid, newLine)
		i++
	}
	return fmt.Sprint(grid)
}