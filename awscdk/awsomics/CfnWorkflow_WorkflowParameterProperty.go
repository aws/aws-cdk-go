package awsomics


// A workflow parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowParameterProperty := &WorkflowParameterProperty{
//   	Description: jsii.String("description"),
//   	Optional: jsii.Boolean(false),
//   }
//
type CfnWorkflow_WorkflowParameterProperty struct {
	// The parameter's description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the parameter is optional.
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

