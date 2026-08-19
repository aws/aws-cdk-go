package awscdkgluealpha


// Properties for configuring an on-demand Glue Trigger.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var stack Stack
//   var role IRole
//   var script Code
//
//
//   // Create a job to run from the workflow
//   job := glue.NewPySparkEtlJob(stack, jsii.String("Job"), &PySparkEtlJobProps{
//   	Role: Role,
//   	Script: Script,
//   })
//
//   // Create a workflow and add a trigger that runs the job
//   workflow := glue.NewWorkflow(stack, jsii.String("Workflow"))
//   workflow.AddOnDemandTrigger(jsii.String("OnDemandTrigger"), &OnDemandTriggerOptions{
//   	Actions: []Action{
//   		&Action{
//   			Job: *Job,
//   		},
//   	},
//   })
//
// Experimental.
type OnDemandTriggerOptions struct {
	// The actions initiated by this trigger.
	// Experimental.
	Actions *[]*Action `field:"required" json:"actions" yaml:"actions"`
	// A description for the trigger.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the trigger.
	// Default: - no name is provided.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

