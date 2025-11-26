package awsimagebuilderalpha


// Configuration details for a workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   var workflow Workflow
//   var workflowParameterValue WorkflowParameterValue
//
//   workflowConfiguration := &WorkflowConfiguration{
//   	Workflow: workflow,
//
//   	// the properties below are optional
//   	OnFailure: imagebuilder_alpha.WorkflowOnFailure_ABORT,
//   	ParallelGroup: jsii.String("parallelGroup"),
//   	Parameters: map[string]WorkflowParameterValue{
//   		"parametersKey": workflowParameterValue,
//   	},
//   }
//
// Experimental.
type WorkflowConfiguration struct {
	// The workflow to execute in the image build.
	// Experimental.
	Workflow IWorkflow `field:"required" json:"workflow" yaml:"workflow"`
	// The action to take if the workflow fails.
	// Default: WorkflowOnFailure.ABORT
	//
	// Experimental.
	OnFailure WorkflowOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The named parallel group to include this workflow in.
	//
	// Workflows in the same parallel group run in parallel of each
	// other.
	// Default: None.
	//
	// Experimental.
	ParallelGroup *string `field:"optional" json:"parallelGroup" yaml:"parallelGroup"`
	// The parameters to pass to the workflow at execution time.
	// Default: - none if the workflow has no parameters, otherwise the default parameter values are used.
	//
	// Experimental.
	Parameters *map[string]WorkflowParameterValue `field:"optional" json:"parameters" yaml:"parameters"`
}

