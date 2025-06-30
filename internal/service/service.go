package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewTeacherServiceService,
	NewClassServiceService,
	NewStudentServiceService,
	NewHistoryService,
	NewHistoryConsumer,
	ProvideKafkaBrokers,
	ProvideKafkaGroupID,
	ProvideKafkaTopic,
)

// ProvideKafkaBrokers returns the list of Kafka broker addresses used by the application.
func ProvideKafkaBrokers() KafkaBrokers {
	return KafkaBrokers([]string{"localhost:9092"})
}

// ProvideKafkaGroupID returns the Kafka consumer group ID for the history service.
func ProvideKafkaGroupID() KafkaGroupID {
	return KafkaGroupID("history-consumer-group")
}

// ProvideKafkaTopic returns the Kafka topic name used for consuming history-related messages.
func ProvideKafkaTopic() KafkaTopic {
	return KafkaTopic("history-topic")
}
