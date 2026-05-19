package awsbedrockagentcore


// Configuration for filtering API Gateway tools.
//
// Tool filters allow you to select REST API operations using path and method combinations.
// Each filter supports two path matching strategies:
// - **Explicit paths**: Matches a single specific path, such as `/pets/{petId}`
// - **Wildcard paths**: Matches all paths starting with the specified prefix, such as `/pets/*`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiGatewayHttpMethod ApiGatewayHttpMethod
//
//   apiGatewayToolFilter := &ApiGatewayToolFilter{
//   	FilterPath: jsii.String("filterPath"),
//   	Methods: []ApiGatewayHttpMethod{
//   		apiGatewayHttpMethod,
//   	},
//   }
//
type ApiGatewayToolFilter struct {
	// The resource path to filter Can be an explicit path (e.g., `/pets/{petId}`) or a wildcard path (e.g., `/pets/*`) Must start with a forward slash.
	FilterPath *string `field:"required" json:"filterPath" yaml:"filterPath"`
	// List of HTTP methods to include for this path Each filter specifies both a path and a list of HTTP methods Multiple filters can overlap and duplicates are automatically de-duplicated.
	Methods *[]ApiGatewayHttpMethod `field:"required" json:"methods" yaml:"methods"`
}

