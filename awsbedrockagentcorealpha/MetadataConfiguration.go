package awsbedrockagentcorealpha


// Configuration for passing metadata (headers and query parameters) to the API Gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   metadataConfiguration := &MetadataConfiguration{
//   	AllowedQueryParameters: []*string{
//   		jsii.String("allowedQueryParameters"),
//   	},
//   	AllowedRequestHeaders: []*string{
//   		jsii.String("allowedRequestHeaders"),
//   	},
//   	AllowedResponseHeaders: []*string{
//   		jsii.String("allowedResponseHeaders"),
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type MetadataConfiguration struct {
	// List of query parameter names to pass through to the target.
	//
	// Constraints:
	// - Array must contain 1-10 items
	// - Each parameter name must be 1-40 characters
	// - Cannot be an empty array.
	// Default: - No query parameters are passed through.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AllowedQueryParameters *[]*string `field:"optional" json:"allowedQueryParameters" yaml:"allowedQueryParameters"`
	// List of request header names to pass through to the target.
	//
	// Constraints:
	// - Array must contain 1-10 items
	// - Each header name must be 1-100 characters
	// - Cannot be an empty array.
	// Default: - No request headers are passed through.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AllowedRequestHeaders *[]*string `field:"optional" json:"allowedRequestHeaders" yaml:"allowedRequestHeaders"`
	// List of response header names to pass through from the target.
	//
	// Constraints:
	// - Array must contain 1-10 items
	// - Each header name must be 1-100 characters
	// - Cannot be an empty array.
	// Default: - No response headers are passed through.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
}

