package awsrds


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessV2ScalingConfigurationProperty := &serverlessV2ScalingConfigurationProperty{
//   	maxCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   }
//
type CfnDBCluster_ServerlessV2ScalingConfigurationProperty struct {
	// `CfnDBCluster.ServerlessV2ScalingConfigurationProperty.MaxCapacity`.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// `CfnDBCluster.ServerlessV2ScalingConfigurationProperty.MinCapacity`.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

