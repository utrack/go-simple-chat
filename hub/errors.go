package hub

import (
	"errors"
)

// ErrNickCollision is returned by Hub.RegisterClient()
// if connecting client wants to have the nickname
// that is already occupied.
var ErrNickCollision = errors.New("Nickname is occupied! Please choose another one...")
