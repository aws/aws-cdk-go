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
//   // Enable basic event and error metrics
//   myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
//   	ClusterArn: jsii.String(ClusterArn),
//   	Topic: jsii.String("basic-monitoring"),
//   	StartingPosition: lambda.StartingPosition_LATEST,
//   	// Provisioned mode is required for observability features
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MinimumPollers: jsii.Number(2),
//   		MaximumPollers: jsii.Number(10),
//   	},
//   	MetricsConfig: &MetricsConfig{
//   		Metrics: []MetricType{
//   			lambda.MetricType_EVENT_COUNT,
//   			lambda.MetricType_ERROR_COUNT,
//   		},
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

