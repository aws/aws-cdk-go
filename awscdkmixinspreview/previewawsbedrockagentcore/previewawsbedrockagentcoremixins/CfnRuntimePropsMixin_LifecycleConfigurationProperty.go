package previewawsbedrockagentcoremixins


// Configuration for managing the lifecycle of runtime sessions and resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lifecycleConfigurationProperty := &LifecycleConfigurationProperty{
//   	IdleRuntimeSessionTimeout: jsii.Number(123),
//   	MaxLifetime: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-lifecycleconfiguration.html
//
type CfnRuntimePropsMixin_LifecycleConfigurationProperty struct {
	// Timeout in seconds for idle runtime sessions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-lifecycleconfiguration.html#cfn-bedrockagentcore-runtime-lifecycleconfiguration-idleruntimesessiontimeout
	//
	IdleRuntimeSessionTimeout *float64 `field:"optional" json:"idleRuntimeSessionTimeout" yaml:"idleRuntimeSessionTimeout"`
	// Maximum lifetime in seconds for runtime sessions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-lifecycleconfiguration.html#cfn-bedrockagentcore-runtime-lifecycleconfiguration-maxlifetime
	//
	MaxLifetime *float64 `field:"optional" json:"maxLifetime" yaml:"maxLifetime"`
}

