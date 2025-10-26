package subscribes

import (
	"encoding/json"
	"log"
	"time"
	"wb_labs_l0/backend/internal/database"
	"wb_labs_l0/backend/internal/model"

	"github.com/nats-io/nats.go"
)

func MakeJSONRequest(nc *nats.Conn) {
	requestJSON, err := json.MarshalIndent(database.GetOrderByID(1), "", "  ")
	if err != nil {
		log.Printf("Error marshaling request: %v", err)
		return
	}

	response, err := nc.Request("nats-streaming.*", requestJSON, 5*time.Second)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return
	}

	// Парсим ответ
	var responseData struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Order   model.Order `json:"order"`
	}

	if err := json.Unmarshal(response.Data, &responseData); err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		return
	}

	if responseData.Success {
		log.Println("Success")
	} else {
		log.Printf("Error: %s", responseData.Message)
	}
}
