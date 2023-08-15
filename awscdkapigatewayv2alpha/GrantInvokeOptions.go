package awscdkapigatewayv2alpha


// Options for granting invoke access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   grantInvokeOptions := &GrantInvokeOptions{
//   	HttpMethods: []httpMethod{
//   		apigatewayv2_alpha.*httpMethod_ANY,
//   	},
//   }
//
// Experimental.
type GrantInvokeOptions struct {
	// The HTTP methods to allow.
	// Default: - the HttpMethod of the route.
	//
	// Experimental.
	HttpMethods *[]HttpMethod `field:"optional" json:"httpMethods" yaml:"httpMethods"`
}

