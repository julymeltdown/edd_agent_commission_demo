package consumer

import (
	"context"
	event2 "edd_agent_commission/internal/event"
	"edd_agent_commission/internal/service"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func StartConsumer(ctx context.Context, brokers []string, topic string, service *service.CommissionService) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "commission-group",
	})

	for {

		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Could not read message:", err)
			break // 또는 적절한 에러 처리
		}

		var event event2.InsuranceApplicationAcceptedEvent
		err = json.Unmarshal(m.Value, &event)
		if err != nil {
			fmt.Println("Could not unmarshal message:", err)
			continue // 다음 메시지 처리
		}

		err = service.ProcessCommission(ctx, event)
		if err != nil {
			fmt.Println("Could not process commission:", err)
			// 에러 처리 - 재시도, 로깅 등
		}
	}

	if err := r.Close(); err != nil {
		fmt.Println("Could not close reader:", err)
	}

}
