package storage

type Message struct {
	ID 			int 		`json:"id"`
	Text 		string 		`json:"text"`
	CreatedAt 	string  	`json:"created_at"`
}


var StorageMessage []Message