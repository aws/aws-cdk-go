package interfacesawsapigatewayv2


// A reference to a Integration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationReference := &IntegrationReference{
//   	ApiId: jsii.String("apiId"),
//   	IntegrationId: jsii.String("integrationId"),
//   }
//
type IntegrationReference struct {
	// The ApiId of the Integration resource.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The IntegrationId of the Integration resource.
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
}

