package main

type contact struct {
	sendingLimit uint8
	userID       string
	age          uint8
}

type perms struct {
	canSend         bool
	canReceive      bool
	permissionLevel uint8
	canManage       bool
}
