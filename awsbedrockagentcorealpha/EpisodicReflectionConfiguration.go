package awsbedrockagentcorealpha


// Configuration for episodic memory reflection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   episodicReflectionConfiguration := &EpisodicReflectionConfiguration{
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type EpisodicReflectionConfiguration struct {
	// Namespaces for episodic reflection Minimum 1 namespace required.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
}

