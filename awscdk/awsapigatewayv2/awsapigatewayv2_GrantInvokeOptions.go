package awsapigatewayv2


// Options for granting invoke access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grantInvokeOptions := &grantInvokeOptions{
//   	httpMethods: []httpMethod{
//   		awscdk.Aws_apigatewayv2.*httpMethod_ANY,
//   	},
//   }
//
// Experimental.
type GrantInvokeOptions struct {
	// The HTTP methods to allow.
	// Experimental.
	HttpMethods *[]HttpMethod `field:"optional" json:"httpMethods" yaml:"httpMethods"`
}

