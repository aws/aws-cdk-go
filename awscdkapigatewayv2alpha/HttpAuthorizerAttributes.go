package awscdkapigatewayv2alpha


// Reference to an http authorizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   httpAuthorizerAttributes := &HttpAuthorizerAttributes{
//   	AuthorizerId: jsii.String("authorizerId"),
//   	AuthorizerType: jsii.String("authorizerType"),
//   }
//
// Deprecated.
type HttpAuthorizerAttributes struct {
	// Id of the Authorizer.
	// Deprecated.
	AuthorizerId *string `field:"required" json:"authorizerId" yaml:"authorizerId"`
	// Type of authorizer.
	//
	// Possible values are:
	// - JWT - JSON Web Token Authorizer
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization.
	// Deprecated.
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
}

