package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
)

// An action initiated by a trigger.
//
// An action runs exactly one target: use {@link Action.job} to run a job or
// {@link Action.crawler} to run a crawler. Because these are separate factory
// methods, an action can never target both or neither.
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
//   		glue.Action_Job(job),
//   	},
//   })
//
// Experimental.
type Action interface {
}

// The jsii proxy struct for Action
type jsiiProxy_Action struct {
	_ byte // padding
}

// Experimental.
func NewAction_Override(a Action) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Action",
		nil, // no parameters
		a,
	)
}

// Create an action that runs a crawler.
// Experimental.
func Action_Crawler(crawler interfacesawsglue.ICrawlerRef, options *CrawlerActionOptions) Action {
	_init_.Initialize()

	if err := validateAction_CrawlerParameters(crawler, options); err != nil {
		panic(err)
	}
	var returns Action

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Action",
		"crawler",
		[]interface{}{crawler, options},
		&returns,
	)

	return returns
}

// Create an action that runs a job.
// Experimental.
func Action_Job(job interfacesawsglue.IJobRef, options *JobActionOptions) Action {
	_init_.Initialize()

	if err := validateAction_JobParameters(job, options); err != nil {
		panic(err)
	}
	var returns Action

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Action",
		"job",
		[]interface{}{job, options},
		&returns,
	)

	return returns
}

