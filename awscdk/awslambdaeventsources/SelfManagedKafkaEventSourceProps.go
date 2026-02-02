package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for a self managed Kafka cluster event source.
//
// If your Kafka cluster is only reachable via VPC make sure to configure it.
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
	//   * 1000 for `DynamoEventSource`
	// * 10000 for `KinesisEventSource`, `ManagedKafkaEventSource` and `SelfManagedKafkaEventSource`.
	// Default: 100.
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5).
	// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-batching
	//
	// Default: - Duration.seconds(0) for Kinesis, DynamoDB, and SQS event sources, Duration.millis(500) for MSK, self-managed Kafka, and Amazon MQ.
	//
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// Configuration for provisioned pollers that read from the event source.
	//
	// When specified, allows control over the minimum and maximum number of pollers
	// that can be provisioned to process events from the source.
	// Default: - no provisioned pollers.
	//
	ProvisionedPollerConfig *ProvisionedPollerConfig `field:"optional" json:"provisionedPollerConfig" yaml:"provisionedPollerConfig"`
	// The Kafka topic to subscribe to.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// * If the function returns an error, split the batch in two and retry.
	// Default: false.
	//
	BisectBatchOnError *bool `field:"optional" json:"bisectBatchOnError" yaml:"bisectBatchOnError"`
	// The identifier for the Kafka consumer group to join.
	//
	// The consumer group ID must be unique among all your Kafka event sources. After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value.  The value must have a length between 1 and 200 and full the pattern '[a-zA-Z0-9-\/*:_+=.@-]*'.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-consumer-group-id
	//
	// Default: - none.
	//
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// Add Customer managed KMS key to encrypt Filter Criteria.
	// See: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
	//
	// Default: - none.
	//
	FilterEncryption awskms.IKey `field:"optional" json:"filterEncryption" yaml:"filterEncryption"`
	// Add filter criteria to Event Source.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html
	//
	// Default: - none.
	//
	Filters *[]*map[string]interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Configuration for logging verbosity from the event source mapping poller.
	// Default: - No logging.
	//
	LogLevel awslambda.EventSourceMappingLogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// The default value is -1, which sets the maximum age to infinite.
	// When the value is set to infinite, Lambda never discards old records.
	// Record are valid until it expires in the event source.
	// Default: -1.
	//
	MaxRecordAge awscdk.Duration `field:"optional" json:"maxRecordAge" yaml:"maxRecordAge"`
	// Configuration for enhanced monitoring metrics collection.
	// Default: - Enhanced monitoring is disabled.
	//
	MetricsConfig *awslambda.MetricsConfig `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// Add an on Failure Destination for this Kafka event.
	//
	// Supported destinations:
	// - {@link KafkaDlq } - Send failed records to a Kafka topic
	// - SNS topics - Send failed records to an SNS topic
	// - SQS queues - Send failed records to an SQS queue
	// - S3 buckets - Send failed records to an S3 bucket.
	// Default: - discarded records are ignored.
	//
	OnFailure awslambda.IEventSourceDlq `field:"optional" json:"onFailure" yaml:"onFailure"`
	// * Allow functions to return partially successful responses for a batch of records.
	// Default: false.
	//
	ReportBatchItemFailures *bool `field:"optional" json:"reportBatchItemFailures" yaml:"reportBatchItemFailures"`
	// * Maximum number of retry attempts.
	//
	// Set to -1 for infinite retries (until the record expires in the event source).
	// Default: -1 (infinite retries).
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Specific configuration settings for a Kafka schema registry.
	// Default: - none.
	//
	SchemaRegistryConfig awslambda.ISchemaRegistry `field:"optional" json:"schemaRegistryConfig" yaml:"schemaRegistryConfig"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	// Default: none.
	//
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// The time from which to start reading, in Unix time seconds.
	// Default: - no timestamp.
	//
	StartingPositionTimestamp *float64 `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
	// The list of host and port pairs that are the addresses of the Kafka brokers in a "bootstrap" Kafka cluster that a Kafka client connects to initially to bootstrap itself.
	//
	// They are in the format `abc.xyz.com:xxxx`.
	BootstrapServers *[]*string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// The authentication method for your Kafka cluster.
	// Default: AuthenticationMethod.SASL_SCRAM_512_AUTH
	//
	AuthenticationMethod AuthenticationMethod `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// The secret with the root CA certificate used by your Kafka brokers for TLS encryption This field is required if your Kafka brokers use certificates signed by a private CA.
	// Default: - none.
	//
	RootCACertificate awssecretsmanager.ISecret `field:"optional" json:"rootCACertificate" yaml:"rootCACertificate"`
	// If your Kafka brokers are only reachable via VPC, provide the security group here.
	// Default: - none, required if setting vpc.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// If your Kafka brokers are only reachable via VPC provide the VPC here.
	// Default: none.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// If your Kafka brokers are only reachable via VPC, provide the subnets selection here.
	// Default: - none, required if setting vpc.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

