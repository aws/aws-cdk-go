package awslambdaeventsources


// The authentication method to use with SelfManagedKafkaEventSource.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // With provisioned pollers and poller group for cost optimization
//   var myFunction Function
//   var kafkaCredentials ISecret
//
//   myFunction.AddEventSource(awscdk.NewSelfManagedKafkaEventSource(&SelfManagedKafkaEventSourceProps{
//   	BootstrapServers: []*string{
//   		jsii.String("kafka-broker1.example.com:9092"),
//   		jsii.String("kafka-broker2.example.com:9092"),
//   	},
//   	Topic: jsii.String("events-topic"),
//   	Secret: kafkaCredentials,
//   	StartingPosition: awscdk.StartingPosition_LATEST,
//   	AuthenticationMethod: awscdk.AuthenticationMethod_SASL_SCRAM_512_AUTH,
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MinimumPollers: jsii.Number(1),
//   		MaximumPollers: jsii.Number(8),
//   		PollerGroupName: jsii.String("self-managed-kafka-group"),
//   	},
//   }))
//
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

