package interfacesawssagemaker


// A reference to a ModelQualityJobDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelQualityJobDefinitionReference := &ModelQualityJobDefinitionReference{
//   	JobDefinitionArn: jsii.String("jobDefinitionArn"),
//   }
//
type ModelQualityJobDefinitionReference struct {
	// The JobDefinitionArn of the ModelQualityJobDefinition resource.
	JobDefinitionArn *string `field:"required" json:"jobDefinitionArn" yaml:"jobDefinitionArn"`
}

