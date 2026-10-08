package interfacesawsobservabilityadmin


// A reference to a DatasetIntegration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetIntegrationReference := &DatasetIntegrationReference{
//   	DatasetIntegrationArn: jsii.String("datasetIntegrationArn"),
//   }
//
type DatasetIntegrationReference struct {
	// The Arn of the DatasetIntegration resource.
	DatasetIntegrationArn *string `field:"required" json:"datasetIntegrationArn" yaml:"datasetIntegrationArn"`
}

