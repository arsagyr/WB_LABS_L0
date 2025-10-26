package subscribes

import (
	"encoding/json"
	"log"
	"wb_labs_l0/backend/internal/database"
	"wb_labs_l0/backend/internal/model"

	"github.com/nats-io/nats.go"
)

func SubscribeToOrderJSON(nc *nats.Conn) {

	_, err := nc.Subscribe("nats-streaming.*", func(msg *nats.Msg) {
		var request model.Order

		if err := json.Unmarshal(msg.Data, &request); err != nil {
			errorResponse := err.Error()
			errorJSON, _ := json.Marshal(errorResponse)
			msg.Respond(errorJSON)
			return
		}
		order := model.Order{
			Order_uid:          request.Order_uid,
			Track_number:       request.Track_number,
			Entry:              request.Entry,
			Delivery:           request.Delivery,
			Payment:            request.Payment,
			Items:              request.Items,
			Locale:             request.Locale,
			Internal_signature: request.Internal_signature,
			Customer_id:        request.Customer_id,
			Delivery_service:   request.Delivery_service,
			Shardkey:           request.Shardkey,
			Sm_id:              request.Sm_id,
			Date_created:       request.Date_created,
			Oof_shard:          request.Oof_shard,
		}

		responseData := order.Order_uid + " is found"

		responseJSON, _ := json.Marshal(responseData)
		msg.Respond(responseJSON)

		log.Printf("Processed request for user ID: %v", request.Order_uid)
		order.Print()
		database.InsertOrder(order)

	})

	if err != nil {
		log.Fatal(err)
	}

}
