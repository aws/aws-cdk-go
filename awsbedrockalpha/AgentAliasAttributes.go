package awsbedrockalpha


// Attributes needed to create an import.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   var agent agent
//
//   agentAliasAttributes := &AgentAliasAttributes{
//   	Agent: agent,
//   	AgentVersion: jsii.String("agentVersion"),
//   	AliasId: jsii.String("aliasId"),
//
//   	// the properties below are optional
//   	AliasName: jsii.String("aliasName"),
//   }
//
// Experimental.
type AgentAliasAttributes struct {
	// The underlying agent for this alias.
	// Experimental.
	Agent IAgent `field:"required" json:"agent" yaml:"agent"`
	// The agent version for this alias.
	// Experimental.
	AgentVersion *string `field:"required" json:"agentVersion" yaml:"agentVersion"`
	// The Id of the agent alias.
	// Experimental.
	AliasId *string `field:"required" json:"aliasId" yaml:"aliasId"`
	// The name of the agent alias.
	// Default: undefined - No alias name is provided.
	//
	// Experimental.
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
}

