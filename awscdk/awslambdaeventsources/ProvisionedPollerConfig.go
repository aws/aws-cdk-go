package awslambdaeventsources


// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // Your MSK cluster arn
//   var clusterArn string
//
//   var myFunction function
//
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//   myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
//   	ClusterArn: jsii.String(ClusterArn),
//   	Topic: jsii.String(Topic),
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MinimumPollers: jsii.Number(1),
//   		MaximumPollers: jsii.Number(3),
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
}

