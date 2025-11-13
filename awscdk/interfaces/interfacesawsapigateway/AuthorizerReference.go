package interfacesawsapigateway


// A reference to a Authorizer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizerReference := &AuthorizerReference{
//   	AuthorizerId: jsii.String("authorizerId"),
//   	RestApiId: jsii.String("restApiId"),
//   }
//
type AuthorizerReference struct {
	// The AuthorizerId of the Authorizer resource.
	AuthorizerId *string `field:"required" json:"authorizerId" yaml:"authorizerId"`
	// The RestApiId of the Authorizer resource.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

