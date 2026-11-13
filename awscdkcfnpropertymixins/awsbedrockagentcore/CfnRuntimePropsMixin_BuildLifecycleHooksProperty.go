package awsbedrockagentcore


// Hooks that run while the platform builds a runtime version, before any session exists.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   buildLifecycleHooksProperty := &BuildLifecycleHooksProperty{
//   	Warmup: &LifecycleHookProperty{
//   		Action: &LifecycleHookActionProperty{
//   			HttpPost: &HttpPostLifecycleHookActionProperty{
//   				Path: jsii.String("path"),
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-buildlifecyclehooks.html
//
type CfnRuntimePropsMixin_BuildLifecycleHooksProperty struct {
	// A lifecycle hook and the action it runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-buildlifecyclehooks.html#cfn-bedrockagentcore-runtime-buildlifecyclehooks-warmup
	//
	Warmup interface{} `field:"optional" json:"warmup" yaml:"warmup"`
}

