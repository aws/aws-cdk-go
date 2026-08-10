package interfacesawsimagebuilder


// A reference to a WorkflowStepExecution resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowStepExecutionReference := &WorkflowStepExecutionReference{
//   	StepExecutionId: jsii.String("stepExecutionId"),
//   }
//
type WorkflowStepExecutionReference struct {
	// The StepExecutionId of the WorkflowStepExecution resource.
	StepExecutionId *string `field:"required" json:"stepExecutionId" yaml:"stepExecutionId"`
}

