package awscdkgluealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2/internal"
)

// The base interface for Glue Workflow.
// See: https://docs.aws.amazon.com/glue/latest/dg/workflows_overview.html
//
// Experimental.
type IWorkflow interface {
	awscdk.IResource
	// Add an custom-scheduled trigger to the workflow.
	// Experimental.
	AddCustomScheduledTrigger(id *string, options *CustomScheduledTriggerOptions) awsglue.CfnTrigger
	// Add an daily-scheduled trigger to the workflow.
	// Experimental.
	AddDailyScheduledTrigger(id *string, options *DailyScheduleTriggerOptions) awsglue.CfnTrigger
	// Add an on-demand trigger to the workflow.
	// Experimental.
	AddOnDemandTrigger(id *string, options *OnDemandTriggerOptions) awsglue.CfnTrigger
	// Add an weekly-scheduled trigger to the workflow.
	// Experimental.
	AddWeeklyScheduledTrigger(id *string, options *WeeklyScheduleTriggerOptions) awsglue.CfnTrigger
	// The ARN of the workflow.
	// Experimental.
	WorkflowArn() *string
	// The name of the workflow.
	// Experimental.
	WorkflowName() *string
}

// The jsii proxy for IWorkflow
type jsiiProxy_IWorkflow struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IWorkflow) AddCustomScheduledTrigger(id *string, options *CustomScheduledTriggerOptions) awsglue.CfnTrigger {
	if err := i.validateAddCustomScheduledTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsglue.CfnTrigger

	_jsii_.Invoke(
		i,
		"addCustomScheduledTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkflow) AddDailyScheduledTrigger(id *string, options *DailyScheduleTriggerOptions) awsglue.CfnTrigger {
	if err := i.validateAddDailyScheduledTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsglue.CfnTrigger

	_jsii_.Invoke(
		i,
		"addDailyScheduledTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkflow) AddOnDemandTrigger(id *string, options *OnDemandTriggerOptions) awsglue.CfnTrigger {
	if err := i.validateAddOnDemandTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsglue.CfnTrigger

	_jsii_.Invoke(
		i,
		"addOnDemandTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkflow) AddWeeklyScheduledTrigger(id *string, options *WeeklyScheduleTriggerOptions) awsglue.CfnTrigger {
	if err := i.validateAddWeeklyScheduledTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsglue.CfnTrigger

	_jsii_.Invoke(
		i,
		"addWeeklyScheduledTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IWorkflow) WorkflowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkflow) WorkflowName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowName",
		&returns,
	)
	return returns
}

