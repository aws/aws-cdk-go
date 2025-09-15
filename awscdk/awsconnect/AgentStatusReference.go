package awsconnect


// A reference to a AgentStatus resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentStatusReference := &AgentStatusReference{
//   	AgentStatusArn: jsii.String("agentStatusArn"),
//   }
//
type AgentStatusReference struct {
	// The AgentStatusArn of the AgentStatus resource.
	AgentStatusArn *string `field:"required" json:"agentStatusArn" yaml:"agentStatusArn"`
}

