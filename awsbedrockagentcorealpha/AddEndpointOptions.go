package awsbedrockagentcorealpha


// Options for adding an endpoint to the runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   addEndpointOptions := &AddEndpointOptions{
//   	Description: jsii.String("description"),
//   	Version: jsii.String("version"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type AddEndpointOptions struct {
	// Description for the endpoint, Must be between 1 and 1200 characters.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Override the runtime version for this endpoint.
	// Default: 1.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

