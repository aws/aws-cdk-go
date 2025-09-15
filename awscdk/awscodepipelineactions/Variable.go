package awscodepipelineactions


// A pipeline-level variable used for a pipeline execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variable := &Variable{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
type Variable struct {
	// The name of a pipeline-level variable.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of a pipeline-level variable.
	Value *string `field:"required" json:"value" yaml:"value"`
}

