package awsapigatewayv2


// Options for granting invoke access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grantInvokeOptions := &GrantInvokeOptions{
//   	HttpMethods: []HttpMethod{
//   		awscdk.Aws_apigatewayv2.HttpMethod_ANY,
//   	},
//   }
//
type GrantInvokeOptions struct {
	// The HTTP methods to allow.
	// Default: - the HttpMethod of the route.
	//
	HttpMethods *[]HttpMethod `field:"optional" json:"httpMethods" yaml:"httpMethods"`
}

