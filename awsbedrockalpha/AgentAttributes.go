package awsbedrockalpha


// Attributes for specifying an imported Bedrock Agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   agentAttributes := &AgentAttributes{
//   	AgentArn: jsii.String("agentArn"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AgentVersion: jsii.String("agentVersion"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	LastUpdated: jsii.String("lastUpdated"),
//   }
//
// Experimental.
type AgentAttributes struct {
	// The ARN of the agent.
	// Experimental.
	AgentArn *string `field:"required" json:"agentArn" yaml:"agentArn"`
	// The ARN of the IAM role associated to the agent.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The agent version.
	//
	// If no explicit versions have been created,
	// leave this empty to use the DRAFT version. Otherwise, use the
	// version number (e.g. 1).
	// Default: 'DRAFT'.
	//
	// Experimental.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Optional KMS encryption key associated with this agent.
	// Default: undefined - An AWS managed key is used.
	//
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// When this agent was last updated.
	// Default: undefined - No last updated timestamp is provided.
	//
	// Experimental.
	LastUpdated *string `field:"optional" json:"lastUpdated" yaml:"lastUpdated"`
}

