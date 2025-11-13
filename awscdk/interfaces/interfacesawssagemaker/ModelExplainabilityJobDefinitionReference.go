package interfacesawssagemaker


// A reference to a ModelExplainabilityJobDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelExplainabilityJobDefinitionReference := &ModelExplainabilityJobDefinitionReference{
//   	JobDefinitionArn: jsii.String("jobDefinitionArn"),
//   }
//
type ModelExplainabilityJobDefinitionReference struct {
	// The JobDefinitionArn of the ModelExplainabilityJobDefinition resource.
	JobDefinitionArn *string `field:"required" json:"jobDefinitionArn" yaml:"jobDefinitionArn"`
}

