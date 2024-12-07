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
//   }
//
type ProvisionedPollerConfig struct {
	// The maximum number of pollers that can be provisioned.
	// Default: - 200.
	//
	MaximumPollers *float64 `field:"optional" json:"maximumPollers" yaml:"maximumPollers"`
	// The minimum number of pollers that should be provisioned.
	// Default: - 1.
	//
	MinimumPollers *float64 `field:"optional" json:"minimumPollers" yaml:"minimumPollers"`
}

