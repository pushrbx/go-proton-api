package server

import "github.com/pushrbx/go-proton-api/server/backend"

func init() {
	backend.GenerateKey = backend.FastGenerateKey
}
