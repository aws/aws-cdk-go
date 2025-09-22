package awsbedrock


// A reference to a AgentAlias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentAliasReference := &AgentAliasReference{
//   	AgentAliasArn: jsii.String("agentAliasArn"),
//   	AgentAliasId: jsii.String("agentAliasId"),
//   	AgentId: jsii.String("agentId"),
//   }
//
type AgentAliasReference struct {
	// The ARN of the AgentAlias resource.
	AgentAliasArn *string `field:"required" json:"agentAliasArn" yaml:"agentAliasArn"`
	// The AgentAliasId of the AgentAlias resource.
	AgentAliasId *string `field:"required" json:"agentAliasId" yaml:"agentAliasId"`
	// The AgentId of the AgentAlias resource.
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
}

