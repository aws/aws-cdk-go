package interfacesawssagemaker


// A reference to a DataQualityJobDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityJobDefinitionReference := &DataQualityJobDefinitionReference{
//   	JobDefinitionArn: jsii.String("jobDefinitionArn"),
//   }
//
type DataQualityJobDefinitionReference struct {
	// The JobDefinitionArn of the DataQualityJobDefinition resource.
	JobDefinitionArn *string `field:"required" json:"jobDefinitionArn" yaml:"jobDefinitionArn"`
}

