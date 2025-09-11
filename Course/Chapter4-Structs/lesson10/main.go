package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	user := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: 0,
		},
	}

	if user.Type == "premium" {
		user.MessageCharLimit = 1000
	} else {
		user.MessageCharLimit = 100
	}

	return user
}
