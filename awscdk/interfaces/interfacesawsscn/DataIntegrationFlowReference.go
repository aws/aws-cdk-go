package interfacesawsscn


// A reference to a DataIntegrationFlow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataIntegrationFlowReference := &DataIntegrationFlowReference{
//   	DataIntegrationFlowArn: jsii.String("dataIntegrationFlowArn"),
//   }
//
type DataIntegrationFlowReference struct {
	// The Arn of the DataIntegrationFlow resource.
	DataIntegrationFlowArn *string `field:"required" json:"dataIntegrationFlowArn" yaml:"dataIntegrationFlowArn"`
}

