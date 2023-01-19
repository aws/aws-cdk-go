package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Properties for a self managed Kafka cluster event source.
//
// If your Kafka cluster is only reachable via VPC make sure to configure it.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // The secret that allows access to your self hosted Kafka cluster
//   var secret secret
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
//   myFunction.addEventSource(awscdk.NewSelfManagedKafkaEventSource(&selfManagedKafkaEventSourceProps{
//   	bootstrapServers: bootstrapServers,
//   	topic: topic,
//   	secret: secret,
//   	batchSize: jsii.Number(100),
//   	 // default
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   }))
//
// Experimental.
type SelfManagedKafkaEventSourceProps struct {
	// Where to begin consuming the stream.
	// Experimental.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	// * 10000 for {@link KinesisEventSource}, {@link ManagedKafkaEventSource} and {@link SelfManagedKafkaEventSource}.
	// Experimental.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	// Experimental.
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// The Kafka topic to subscribe to.
	// Experimental.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	// Experimental.
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// The list of host and port pairs that are the addresses of the Kafka brokers in a "bootstrap" Kafka cluster that a Kafka client connects to initially to bootstrap itself.
	//
	// They are in the format `abc.xyz.com:xxxx`.
	// Experimental.
	BootstrapServers *[]*string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// The authentication method for your Kafka cluster.
	// Experimental.
	AuthenticationMethod AuthenticationMethod `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// If your Kafka brokers are only reachable via VPC, provide the security group here.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// If your Kafka brokers are only reachable via VPC provide the VPC here.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// If your Kafka brokers are only reachable via VPC, provide the subnets selection here.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

