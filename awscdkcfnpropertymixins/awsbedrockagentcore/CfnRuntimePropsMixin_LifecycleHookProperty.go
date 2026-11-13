package awsbedrockagentcore


// A lifecycle hook and the action it runs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   lifecycleHookProperty := &LifecycleHookProperty{
//   	Action: &LifecycleHookActionProperty{
//   		HttpPost: &HttpPostLifecycleHookActionProperty{
//   			Path: jsii.String("path"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-lifecyclehook.html
//
type CfnRuntimePropsMixin_LifecycleHookProperty struct {
	// The action a lifecycle hook runs.
	//
	// HttpPost is the only action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-lifecyclehook.html#cfn-bedrockagentcore-runtime-lifecyclehook-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
}

