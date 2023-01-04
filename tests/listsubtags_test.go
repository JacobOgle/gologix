package gologix_tests

import (
	"testing"

	"github.com/danomagnum/gologix"
)

func TestSubList(t *testing.T) {

	client := gologix.NewClient("192.168.2.241")
	err := client.Connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		err := client.Disconnect()
		if err != nil {
			t.Errorf("problem disconnecting. %v", err)
		}
	}()

	_, err = client.ListSubTags("Program:gologix_tests", 1)
	if err != nil {
		t.Error(err)
		return
	}

}

func compare_array_order(have, want []int) bool {
	if len(have) != len(want) {
		return false
	}

	for i := range have {
		if have[i] != want[i] {
			return false
		}
	}
	return true
}
