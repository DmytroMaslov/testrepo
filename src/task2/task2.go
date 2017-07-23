//task2
//{"En1":{"S1":4, "S2":5},"En2":{"S1":10, "S2":20}}
package task2
import (
	"errors"
	"encoding/json"
)
const ENVELOPE_COUNT = 2

type Envelope struct {
	S1 float64 `json:"width"`
	S2 float64 `json:"height"`
}

func Run(param []byte) (str string, err error){

	var envs = [ENVELOPE_COUNT]Envelope{*new(Envelope), *new(Envelope)}
	err = json.Unmarshal(param, &envs)
	if err != nil{
		return str, errors.New("can't unmarshal data")
	}
	en1 := envs[0]
	en2 := envs[1]
	err = IsValid(en1, en2)
	if err != nil{
		return str, err
	}

	if en1.isBigest(en2) {

		return "1", nil
	}
	if en2.isBigest(en1){
		return "2", nil
	}
	return "0", nil

}

func (en1 * Envelope) isBigest(en2 Envelope) bool{
	isParallel := en1.S1 > en2.S1 && en1.S2>en2.S2
	isPerpendicularly := en1.S1>en2.S2 && en1.S2>en2.S1
	if isParallel || isPerpendicularly{
		return true
	}
	return false
}

func IsValid (en1 Envelope, en2 Envelope) (error){
	if (en1.S1 <= 0|| en1.S2<=0){
		return errors.New("Incorrect first envelope")
	}
	if (en2.S1<= 0|| en2.S2<= 0){
		return errors.New("Incorrect second envelope")
	}
	return nil
}



