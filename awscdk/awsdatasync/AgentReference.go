package awsdatasync


// A reference to a Agent resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentReference := &AgentReference{
//   	AgentArn: jsii.String("agentArn"),
//   }
//
type AgentReference struct {
	// The AgentArn of the Agent resource.
	AgentArn *string `field:"required" json:"agentArn" yaml:"agentArn"`
}

