package auth

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	t.Run("Login Pass ", func(t *testing.T) {
		t.Parallel()
		u := user{"pedro", "123"}
		approved := Auth(u)

		if approved {
			fmt.Printf("User %s logged", u.name)
		} else {
			t.Errorf("User or Pass with error")
		}
	})

	t.Run("Login Fail ", func(t *testing.T) {
		t.Parallel()
		u := user{"pedroerro", "12345"}
		approved := Auth(u)

		if approved {
			fmt.Printf("User %s logged", u.name)
		} else {
			t.Errorf("User or Pass with error")
		}
	})
}
