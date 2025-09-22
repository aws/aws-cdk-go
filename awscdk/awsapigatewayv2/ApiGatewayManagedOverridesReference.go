package awsapigatewayv2


// A reference to a ApiGatewayManagedOverrides resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiGatewayManagedOverridesReference := &ApiGatewayManagedOverridesReference{
//   	ApiGatewayManagedOverridesId: jsii.String("apiGatewayManagedOverridesId"),
//   }
//
type ApiGatewayManagedOverridesReference struct {
	// The Id of the ApiGatewayManagedOverrides resource.
	ApiGatewayManagedOverridesId *string `field:"required" json:"apiGatewayManagedOverridesId" yaml:"apiGatewayManagedOverridesId"`
}

