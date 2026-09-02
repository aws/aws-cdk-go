package interfacesawsapigatewayv2


// A reference to a ApiGatewayManagedOverrides resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiGatewayManagedOverridesReference := &ApiGatewayManagedOverridesReference{
//   	ApiId: jsii.String("apiId"),
//   }
//
type ApiGatewayManagedOverridesReference struct {
	// The ApiId of the ApiGatewayManagedOverrides resource.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
}

