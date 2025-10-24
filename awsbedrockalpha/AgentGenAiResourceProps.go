package awsbedrockalpha


// Properties for creating an agent GenAI resource configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   var agentAlias AgentAlias
//
//   agentGenAiResourceProps := &AgentGenAiResourceProps{
//   	AgentAlias: agentAlias,
//   }
//
// Experimental.
type AgentGenAiResourceProps struct {
	// The agent alias to use for the GenAI resource.
	// Experimental.
	AgentAlias IAgentAlias `field:"required" json:"agentAlias" yaml:"agentAlias"`
}

