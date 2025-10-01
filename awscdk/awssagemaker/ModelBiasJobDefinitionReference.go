package awssagemaker


// A reference to a ModelBiasJobDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelBiasJobDefinitionReference := &ModelBiasJobDefinitionReference{
//   	JobDefinitionArn: jsii.String("jobDefinitionArn"),
//   }
//
type ModelBiasJobDefinitionReference struct {
	// The JobDefinitionArn of the ModelBiasJobDefinition resource.
	JobDefinitionArn *string `field:"required" json:"jobDefinitionArn" yaml:"jobDefinitionArn"`
}

