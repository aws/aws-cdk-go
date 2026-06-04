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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-lifecycleconfiguration.html#cfn-bedrockagentcore-harness-lifecycleconfiguration-idleruntimesessiontimeout
	//
	IdleRuntimeSessionTimeout *float64 `field:"optional" json:"idleRuntimeSessionTimeout" yaml:"idleRuntimeSessionTimeout"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-lifecycleconfiguration.html#cfn-bedrockagentcore-harness-lifecycleconfiguration-maxlifetime
	//
	MaxLifetime *float64 `field:"optional" json:"maxLifetime" yaml:"maxLifetime"`
}

