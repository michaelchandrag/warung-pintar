package contract

import (
	storage "github.com/michaelchandrag/warung-pintar/module/storage"
)

type MessageContract interface {
	Insert() (storage.Message, error)
	Find() ([]storage.Message, error)
}