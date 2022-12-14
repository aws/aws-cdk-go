package awslambdaeventsources


// The authentication method to use with SelfManagedKafkaEventSource.
type AuthenticationMethod string

const (
	// SASL_SCRAM_512_AUTH authentication method for your Kafka cluster.
	AuthenticationMethod_SASL_SCRAM_512_AUTH AuthenticationMethod = "SASL_SCRAM_512_AUTH"
	// SASL_SCRAM_256_AUTH authentication method for your Kafka cluster.
	AuthenticationMethod_SASL_SCRAM_256_AUTH AuthenticationMethod = "SASL_SCRAM_256_AUTH"
	// BASIC_AUTH (SASL/PLAIN) authentication method for your Kafka cluster.
	AuthenticationMethod_BASIC_AUTH AuthenticationMethod = "BASIC_AUTH"
	// CLIENT_CERTIFICATE_TLS_AUTH (mTLS) authentication method for your Kafka cluster.
	AuthenticationMethod_CLIENT_CERTIFICATE_TLS_AUTH AuthenticationMethod = "CLIENT_CERTIFICATE_TLS_AUTH"
)

