package interfacesawsbedrockagentcore


// A reference to a RuntimeEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeEndpointReference := &RuntimeEndpointReference{
//   	AgentRuntimeEndpointArn: jsii.String("agentRuntimeEndpointArn"),
//   }
//
type RuntimeEndpointReference struct {
	// The AgentRuntimeEndpointArn of the RuntimeEndpoint resource.
	AgentRuntimeEndpointArn *string `field:"required" json:"agentRuntimeEndpointArn" yaml:"agentRuntimeEndpointArn"`
}

