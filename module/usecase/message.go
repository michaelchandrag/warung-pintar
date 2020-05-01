package usecase

import (
	"fmt"

	storage "github.com/michaelchandrag/warung-pintar/module/storage"
	contract "github.com/michaelchandrag/warung-pintar/module/contract"
)

type MessageUsecase struct {
	message storage.Message
}

func (mu *MessageUsecase) Insert (mc contract.MessageContract) (result storage.Message, err error) {
	result, err = mc.Insert()
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, nil
}

func (mu *MessageUsecase) Find (mc contract.MessageContract) (result []storage.Message, err error) {
	result, err = mc.Find()
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, nil
}