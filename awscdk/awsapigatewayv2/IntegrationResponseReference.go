package awsapigatewayv2


// A reference to a IntegrationResponse resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationResponseReference := &IntegrationResponseReference{
//   	ApiId: jsii.String("apiId"),
//   	IntegrationId: jsii.String("integrationId"),
//   	IntegrationResponseId: jsii.String("integrationResponseId"),
//   }
//
type IntegrationResponseReference struct {
	// The ApiId of the IntegrationResponse resource.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The IntegrationId of the IntegrationResponse resource.
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The IntegrationResponseId of the IntegrationResponse resource.
	IntegrationResponseId *string `field:"required" json:"integrationResponseId" yaml:"integrationResponseId"`
}

