package task2

import (
	"testing"
	"strings"
)
//{"S1":4, "S2":5},"En2":{"S1":10, "S2":20}}
//[{width: 8, height: 5}, {width: 6, height: 9}]
func TestGetSmallEnvelope(t *testing.T) {
	var EnvelopTestData = []struct {
		Data []byte
		expectedResult string
	}{
		{[]byte (`[{"width":"5","height":"10"},{"width":"10","height":"20"}]`), "2"},
		{[]byte (`{"En1":{"S1": 50, "S2": 100}, "En2": {"S1":1, "S2":2}}`), "1"},
		{[]byte (`{"En1":{"S1": 5, "S2": 10}, "En2": {"S1":5, "S2":10}}`), "0"},
		{[]byte (`{"En1":{"S1": 50, "S2": 5}, "En2": {"S1":4, "S2":49}}`), "1"},
		{[]byte (`{"En1":{"S1": 5, "S2": 10}, "En2": {"S1":10, "S2":5.1}}`), "0"},
		{[]byte (`{"En1":{"S1": 0.00015, "S2": 0.00020}, "En2": {"S1":0.001, "S2":0.002}}`), "2"},
	}
	for _,en := range EnvelopTestData{
		actualRes,_ := Run(en.Data)
		t.Errorf("as"+actualRes)
		if !strings.EqualFold(actualRes, en.expectedResult) {
			t.Errorf("Input: envelop1 envelop2:  %s expected: %s actual: %s", en.Data, en.expectedResult, actualRes)
		}
	}

}

func TestIsValid(t *testing.T) {
	var envelopTestData = []struct {
		en1 Envelope
		en2 Envelope
		errorMessage string
	}{
		{Envelope{S1: -5, S2: 10}, Envelope{S1:10, S2:20}, "Incorrect first envelope"},
		{Envelope{S1: 0, S2: 10}, Envelope{S1:10, S2:20}, "Incorrect first envelope"},
		{Envelope{S1: 1, S2: 10}, Envelope{S1:-10, S2:20}, "Incorrect second envelope"},
	}
	for _, en := range envelopTestData{
		er := IsValid(en.en1, en.en2)
		if er == nil {
			t.Errorf("Input: envelop1: %v envelop2: %v expected: Error actual: nil",en.en1, en.en2)
		}
		if er !=nil && er.Error() != en.errorMessage{
			t.Errorf("Input: envelop1: %v envelop2: %v expected:%s actual:%s ",en.en1, en.en2, en.errorMessage, er.Error())
		}
	}
}