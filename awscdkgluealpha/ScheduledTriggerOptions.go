package awscdkgluealpha


// Properties for configuring a scheduled Glue Trigger.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var stack Stack
//   var role IRole
//   var script Code
//
//   job := glue.NewPySparkEtlJob(stack, jsii.String("Job"), &PySparkEtlJobProps{
//   	Role: Role,
//   	Script: Script,
//   })
//   workflow := glue.NewWorkflow(stack, jsii.String("Workflow"))
//
//   workflow.AddScheduledTrigger(jsii.String("WeeklyTrigger"), &ScheduledTriggerOptions{
//   	Actions: []Action{
//   		glue.Action_Job(job),
//   	},
//   	Schedule: glue.TriggerSchedule_Weekly(),
//   })
//
// Experimental.
type ScheduledTriggerOptions struct {
	// The actions initiated by this trigger.
	// Experimental.
	Actions *[]Action `field:"required" json:"actions" yaml:"actions"`
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
	// Whether to start the trigger on creation or not.
	// Default: - false.
	//
	// Experimental.
	StartOnCreation *bool `field:"optional" json:"startOnCreation" yaml:"startOnCreation"`
	// The schedule on which this trigger fires.
	//
	// Build one with {@link TriggerSchedule.daily}, {@link TriggerSchedule.weekly},
	// {@link TriggerSchedule.cron}, or {@link TriggerSchedule.expression}.
	// Experimental.
	Schedule TriggerSchedule `field:"required" json:"schedule" yaml:"schedule"`
}

