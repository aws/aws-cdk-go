package awsapigatewayv2


// Reference to an http authorizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpAuthorizerAttributes := &httpAuthorizerAttributes{
//   	authorizerId: jsii.String("authorizerId"),
//   	authorizerType: jsii.String("authorizerType"),
//   }
//
// Experimental.
type HttpAuthorizerAttributes struct {
	// Id of the Authorizer.
	// Experimental.
	AuthorizerId *string `field:"required" json:"authorizerId" yaml:"authorizerId"`
	// Type of authorizer.
	//
	// Possible values are:
	// - JWT - JSON Web Token Authorizer
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization.
	// Experimental.
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
}

