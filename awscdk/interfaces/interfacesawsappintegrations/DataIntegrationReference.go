package interfacesawsappintegrations


// A reference to a DataIntegration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataIntegrationReference := &DataIntegrationReference{
//   	DataIntegrationArn: jsii.String("dataIntegrationArn"),
//   	DataIntegrationId: jsii.String("dataIntegrationId"),
//   }
//
type DataIntegrationReference struct {
	// The ARN of the DataIntegration resource.
	DataIntegrationArn *string `field:"required" json:"dataIntegrationArn" yaml:"dataIntegrationArn"`
	// The Id of the DataIntegration resource.
	DataIntegrationId *string `field:"required" json:"dataIntegrationId" yaml:"dataIntegrationId"`
}

