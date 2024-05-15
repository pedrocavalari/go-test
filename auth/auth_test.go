package auth

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	u := user{"pedro", "123"}
	approved := Auth(u)

	if approved {
		fmt.Printf("User %s logged", u.name)
	} else {
		t.Errorf("User or Pass with error")
	}

}
