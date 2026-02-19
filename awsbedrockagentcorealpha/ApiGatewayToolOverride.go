package awsbedrockagentcorealpha


// Configuration for overriding API Gateway tool metadata.
//
// Tool overrides allow you to customize the tool name or description for specific operations
// after filtering. Each override must specify an explicit path and a single HTTP method.
// The override must match an operation that exists in your API and must correspond to one
// of the operations resolved by your filters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   apiGatewayToolOverride := &ApiGatewayToolOverride{
//   	Method: bedrock_agentcore_alpha.ApiGatewayHttpMethod_GET,
//   	Name: jsii.String("name"),
//   	Path: jsii.String("path"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// Experimental.
type ApiGatewayToolOverride struct {
	// The HTTP method for this override Must be a single method (no wildcards).
	// Experimental.
	Method ApiGatewayHttpMethod `field:"required" json:"method" yaml:"method"`
	// The custom tool name If not provided, the operationId from the OpenAPI definition will be used.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The explicit resource path (no wildcards) Must match an operation that exists in your API.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Optional custom description for the tool.
	// Default: - No custom description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

