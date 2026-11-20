package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleConfigurationProperty := &LifecycleConfigurationProperty{
//   	IdleRuntimeSessionTimeout: jsii.Number(123),
//   	MaxLifetime: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-lifecycleconfiguration.html
//
type CfnHarness_LifecycleConfigurationProperty struct {
	// Timeout in seconds for idle runtime sessions.
	//
	// When a session remains idle for this duration, it will be automatically terminated. Valid range: 60 to 28800 seconds (8 hours), and must be less than or equal to MaxLifetime. Default: 900 seconds (15 minutes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-lifecycleconfiguration.html#cfn-bedrockagentcore-harness-lifecycleconfiguration-idleruntimesessiontimeout
	//
	IdleRuntimeSessionTimeout *float64 `field:"optional" json:"idleRuntimeSessionTimeout" yaml:"idleRuntimeSessionTimeout"`
	// Maximum lifetime for the instance in seconds.
	//
	// Once reached, instances will be automatically terminated and replaced. Valid range: 60 to 28800 seconds (8 hours). Default: 28800 seconds (8 hours).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-lifecycleconfiguration.html#cfn-bedrockagentcore-harness-lifecycleconfiguration-maxlifetime
	//
	MaxLifetime *float64 `field:"optional" json:"maxLifetime" yaml:"maxLifetime"`
}

