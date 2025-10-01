package awsapigatewayv2


// A reference to a Authorizer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizerReference := &AuthorizerReference{
//   	ApiId: jsii.String("apiId"),
//   	AuthorizerId: jsii.String("authorizerId"),
//   }
//
type AuthorizerReference struct {
	// The ApiId of the Authorizer resource.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The AuthorizerId of the Authorizer resource.
	AuthorizerId *string `field:"required" json:"authorizerId" yaml:"authorizerId"`
}

