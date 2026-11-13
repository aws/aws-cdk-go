package awsbedrockagentcore


// Lifecycle hooks that the platform runs at fixed points in the runtime's lifecycle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   runtimeHooksConfigurationProperty := &RuntimeHooksConfigurationProperty{
//   	BuildProperty: &BuildLifecycleHooksProperty{
//   		Warmup: &LifecycleHookProperty{
//   			Action: &LifecycleHookActionProperty{
//   				HttpPost: &HttpPostLifecycleHookActionProperty{
//   					Path: jsii.String("path"),
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-runtimehooksconfiguration.html
//
type CfnRuntimePropsMixin_RuntimeHooksConfigurationProperty struct {
	// Hooks that run while the platform builds a runtime version, before any session exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-runtimehooksconfiguration.html#cfn-bedrockagentcore-runtime-runtimehooksconfiguration-build
	//
	BuildProperty interface{} `field:"optional" json:"buildProperty" yaml:"buildProperty"`
}

