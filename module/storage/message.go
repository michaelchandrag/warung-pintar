package storage

func (this Message) Insert() (result Message, err error) {
	currentID := len(StorageMessage) + 1
	this.ID = currentID
	StorageMessage = append(StorageMessage, this)
	return this, nil
}

func (this Message) Find() (result []Message, err error) {
	return StorageMessage, nil
}