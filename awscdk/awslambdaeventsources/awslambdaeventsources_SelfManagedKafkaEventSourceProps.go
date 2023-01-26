package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for a self managed Kafka cluster event source.
//
// If your Kafka cluster is only reachable via VPC make sure to configure it.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // The secret that allows access to your self hosted Kafka cluster
//   var secret secret
//
//   // (Optional) The secret containing the root CA certificate that your Kafka brokers use for TLS encryption
//   var encryption secret
//
//   var myFunction function
//
//
//   // The list of Kafka brokers
//   bootstrapServers := []*string{
//   	"kafka-broker:9092",
//   }
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//
//   // (Optional) The consumer group id to use when connecting to the Kafka broker. If omitted the UUID of the event source mapping will be used.
//   var consumerGroupId string
//   myFunction.addEventSource(awscdk.NewSelfManagedKafkaEventSource(&selfManagedKafkaEventSourceProps{
//   	bootstrapServers: bootstrapServers,
//   	topic: topic,
//   	consumerGroupId: consumerGroupId,
//   	secret: secret,
//   	batchSize: jsii.Number(100),
//   	 // default
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   	encryption: encryption,
//   }))
//
type SelfManagedKafkaEventSourceProps struct {
	// Where to begin consuming the stream.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for `DynamoEventSource`
	// * 10000 for `KinesisEventSource`, `ManagedKafkaEventSource` and `SelfManagedKafkaEventSource`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5).
	// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-batching
	//
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// The Kafka topic to subscribe to.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The identifier for the Kafka consumer group to join.
	//
	// The consumer group ID must be unique among all your Kafka event sources. After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value.  The value must have a lenght between 1 and 200 and full the pattern '[a-zA-Z0-9-\/*:_+=.@-]*'.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-consumer-group-id
	//
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// The list of host and port pairs that are the addresses of the Kafka brokers in a "bootstrap" Kafka cluster that a Kafka client connects to initially to bootstrap itself.
	//
	// They are in the format `abc.xyz.com:xxxx`.
	BootstrapServers *[]*string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// The authentication method for your Kafka cluster.
	AuthenticationMethod AuthenticationMethod `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// The secret with the root CA certificate used by your Kafka brokers for TLS encryption This field is required if your Kafka brokers use certificates signed by a private CA.
	RootCACertificate awssecretsmanager.ISecret `field:"optional" json:"rootCACertificate" yaml:"rootCACertificate"`
	// If your Kafka brokers are only reachable via VPC, provide the security group here.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// If your Kafka brokers are only reachable via VPC provide the VPC here.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// If your Kafka brokers are only reachable via VPC, provide the subnets selection here.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

