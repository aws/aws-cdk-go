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
// Experimental.
type AddEndpointOptions struct {
	// Description for the endpoint, Must be between 1 and 1200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Override the runtime version for this endpoint.
	// Default: 1.
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

