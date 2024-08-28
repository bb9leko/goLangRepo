package matematica

import "testing"

func TestSoma(t *testing.T) {

	expected := 5

	result := Soma(3, 2)

	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

}
