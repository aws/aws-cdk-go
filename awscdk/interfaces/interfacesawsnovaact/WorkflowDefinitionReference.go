package interfacesawsnovaact


// A reference to a WorkflowDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowDefinitionReference := &WorkflowDefinitionReference{
//   	WorkflowDefinitionArn: jsii.String("workflowDefinitionArn"),
//   }
//
type WorkflowDefinitionReference struct {
	// The Arn of the WorkflowDefinition resource.
	WorkflowDefinitionArn *string `field:"required" json:"workflowDefinitionArn" yaml:"workflowDefinitionArn"`
}

