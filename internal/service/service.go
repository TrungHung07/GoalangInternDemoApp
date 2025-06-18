package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService,
	NewTeacherServiceService,
	NewClassServiceService,
	NewStudentServiceService,
	NewHistoryService,
	NewHistoryConsumer,
	ProvideKafkaBrokers,
	ProvideKafkaGroupID,
	ProvideKafkaTopic,
)

func ProvideKafkaBrokers() KafkaBrokers {
	return KafkaBrokers([]string{"localhost:9092"})
}

func ProvideKafkaGroupID() KafkaGroupID {
	return KafkaGroupID("history-consumer-group")
}

func ProvideKafkaTopic() KafkaTopic {
	return KafkaTopic("history-topic")
}
