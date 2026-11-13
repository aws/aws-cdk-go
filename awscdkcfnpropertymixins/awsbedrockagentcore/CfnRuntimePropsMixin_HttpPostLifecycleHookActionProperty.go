package awsbedrockagentcore


// Sends an HTTP POST to the agent's container.
//
// Any 2xx response ends the hook; anything else fails it. The path and port must not be a route the platform already calls on the container, such as the protocol's invocation path or the default health check.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   httpPostLifecycleHookActionProperty := &HttpPostLifecycleHookActionProperty{
//   	Path: jsii.String("path"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-httppostlifecyclehookaction.html
//
type CfnRuntimePropsMixin_HttpPostLifecycleHookActionProperty struct {
	// Path the platform sends the POST to.
	//
	// Must start with a slash.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-httppostlifecyclehookaction.html#cfn-bedrockagentcore-runtime-httppostlifecyclehookaction-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Port in the agent's container the platform sends the POST to.
	//
	// Defaults to the runtime's server port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-httppostlifecyclehookaction.html#cfn-bedrockagentcore-runtime-httppostlifecyclehookaction-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

