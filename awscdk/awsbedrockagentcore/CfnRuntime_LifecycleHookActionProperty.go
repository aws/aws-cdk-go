package awsbedrockagentcore


// The action a lifecycle hook runs.
//
// HttpPost is the only action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleHookActionProperty := &LifecycleHookActionProperty{
//   	HttpPost: &HttpPostLifecycleHookActionProperty{
//   		Path: jsii.String("path"),
//
//   		// the properties below are optional
//   		Port: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-lifecyclehookaction.html
//
type CfnRuntime_LifecycleHookActionProperty struct {
	// Sends an HTTP POST to the agent's container.
	//
	// Any 2xx response ends the hook; anything else fails it. The path and port must not be a route the platform already calls on the container, such as the protocol's invocation path or the default health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-lifecyclehookaction.html#cfn-bedrockagentcore-runtime-lifecyclehookaction-httppost
	//
	HttpPost interface{} `field:"required" json:"httpPost" yaml:"httpPost"`
}

