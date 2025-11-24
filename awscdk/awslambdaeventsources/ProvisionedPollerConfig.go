package awslambdaeventsources


// (Amazon MSK and self-managed Apache Kafka only) The provisioned mode configuration for the event source.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction Function
//
//
//   // Your MSK cluster arn
//   clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//
//   // Create a Kafka DLQ destination
//   kafkaDlq := awscdk.NewKafkaDlq(jsii.String("failure-topic"))
//
//   myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
//   	ClusterArn: jsii.String(ClusterArn),
//   	Topic: jsii.String(Topic),
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   	OnFailure: kafkaDlq,
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MinimumPollers: jsii.Number(1),
//   		MaximumPollers: jsii.Number(1),
//   	},
//   }))
//
type ProvisionedPollerConfig struct {
	// The maximum number of pollers that can be provisioned.
	// Default: 200.
	//
	MaximumPollers *float64 `field:"required" json:"maximumPollers" yaml:"maximumPollers"`
	// The minimum number of pollers that should be provisioned.
	// Default: 1.
	//
	MinimumPollers *float64 `field:"required" json:"minimumPollers" yaml:"minimumPollers"`
	// An optional identifier that groups multiple ESMs to share EPU capacity and reduce costs.
	//
	// ESMs with the same PollerGroupName share compute
	// resources.
	// Default: - not set, dedicated compute resource per event source.
	//
	PollerGroupName *string `field:"optional" json:"pollerGroupName" yaml:"pollerGroupName"`
}

