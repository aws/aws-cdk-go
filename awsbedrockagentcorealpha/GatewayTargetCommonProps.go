package awsbedrockagentcorealpha


// Common properties for all Gateway Target types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   gatewayTargetCommonProps := &GatewayTargetCommonProps{
//   	Description: jsii.String("description"),
//   	GatewayTargetName: jsii.String("gatewayTargetName"),
//   }
//
// Experimental.
type GatewayTargetCommonProps struct {
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Default: - auto generate.
	//
	// Experimental.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
}

