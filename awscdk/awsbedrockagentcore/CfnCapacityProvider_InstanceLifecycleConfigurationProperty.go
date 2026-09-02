package awsbedrockagentcore


// Configuration for managing the lifecycle of instances in a capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceLifecycleConfigurationProperty := &InstanceLifecycleConfigurationProperty{
//   	IdleInstanceTimeout: jsii.Number(123),
//   	MaxLifetime: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-instancelifecycleconfiguration.html
//
type CfnCapacityProvider_InstanceLifecycleConfigurationProperty struct {
	// The number of seconds an instance can remain idle before it is stopped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-instancelifecycleconfiguration.html#cfn-bedrockagentcore-capacityprovider-instancelifecycleconfiguration-idleinstancetimeout
	//
	IdleInstanceTimeout *float64 `field:"optional" json:"idleInstanceTimeout" yaml:"idleInstanceTimeout"`
	// Maximum lifetime for the instance in seconds.
	//
	// Once reached, instances will be automatically terminated regardless of activity. Default: 28800 seconds (8 hours). Maximum: 1209600 seconds (14 days).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-instancelifecycleconfiguration.html#cfn-bedrockagentcore-capacityprovider-instancelifecycleconfiguration-maxlifetime
	//
	MaxLifetime *float64 `field:"optional" json:"maxLifetime" yaml:"maxLifetime"`
}

