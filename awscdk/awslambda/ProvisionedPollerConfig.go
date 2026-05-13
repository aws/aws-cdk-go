package awslambda


// (Amazon MSK and self-managed Apache Kafka only) The provisioned mode configuration for the event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedPollerConfig := &ProvisionedPollerConfig{
//   	MaximumPollers: jsii.Number(123),
//   	MinimumPollers: jsii.Number(123),
//   	PollerGroupName: jsii.String("pollerGroupName"),
//   }
//
type ProvisionedPollerConfig struct {
	// The maximum number of pollers that can be provisioned.
	//
	// Valid Range:
	// * For Amazon SQS: Minimum value of 2. Maximum value of 2000. Default: 200.
	// * For Amazon MSK, self-managed Apache Kafka, and Amazon MQ: Minimum value of 1. Maximum value of 2000. Default: 200.
	// Default: 200.
	//
	MaximumPollers *float64 `field:"optional" json:"maximumPollers" yaml:"maximumPollers"`
	// The minimum number of pollers that should be provisioned.
	//
	// Valid Range:
	// * For Amazon SQS: Minimum value of 2. Maximum value of 200. Default: 2.
	// * For Amazon MSK, self-managed Apache Kafka, and Amazon MQ: Minimum value of 1. Maximum value of 200. Default: 1.
	// Default: - 2 for SQS, 1 for MSK/Kafka/MQ.
	//
	MinimumPollers *float64 `field:"optional" json:"minimumPollers" yaml:"minimumPollers"`
	// An optional identifier that groups multiple ESMs to share EPU capacity and reduce costs.
	//
	// ESMs with the same PollerGroupName share compute
	// resources.
	// Default: - not set, dedicated compute resource per event source.
	//
	PollerGroupName *string `field:"optional" json:"pollerGroupName" yaml:"pollerGroupName"`
}

