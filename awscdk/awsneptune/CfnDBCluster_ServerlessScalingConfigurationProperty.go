package awsneptune


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessScalingConfigurationProperty := &ServerlessScalingConfigurationProperty{
//   	MaxCapacity: jsii.Number(123),
//   	MinCapacity: jsii.Number(123),
//   }
//
type CfnDBCluster_ServerlessScalingConfigurationProperty struct {
	// `CfnDBCluster.ServerlessScalingConfigurationProperty.MaxCapacity`.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// `CfnDBCluster.ServerlessScalingConfigurationProperty.MinCapacity`.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

