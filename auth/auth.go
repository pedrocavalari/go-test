package auth

type user struct {
	name string
	pass string
}

func Auth(u user) bool {
	u1 := user{"pedro", "123"}
	if u.name == u1.name && u.pass == u1.pass {
		return true
	}
	return false
}
