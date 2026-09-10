package awscdkgluealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2/internal"
)

// The base interface for Glue Workflow.
// See: https://docs.aws.amazon.com/glue/latest/dg/workflows_overview.html
//
// Experimental.
type IWorkflow interface {
	awscdk.IResource
	// Add a conditional (predicate-based) trigger to the workflow.
	//
	// Returns: a reference to the created trigger.
	// Experimental.
	AddConditionalTrigger(id *string, options *ConditionalTriggerOptions) interfacesawsglue.ITriggerRef
	// Add an EventBridge event-based trigger to the workflow.
	//
	// Returns: a reference to the created trigger.
	// Experimental.
	AddEventTrigger(id *string, options *EventTriggerOptions) interfacesawsglue.ITriggerRef
	// Add an on-demand trigger to the workflow.
	//
	// Returns: a reference to the created trigger.
	// Experimental.
	AddOnDemandTrigger(id *string, options *OnDemandTriggerOptions) interfacesawsglue.ITriggerRef
	// Add a scheduled trigger to the workflow.
	//
	// Returns: a reference to the created trigger.
	// Experimental.
	AddScheduledTrigger(id *string, options *ScheduledTriggerOptions) interfacesawsglue.ITriggerRef
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

func (i *jsiiProxy_IWorkflow) AddConditionalTrigger(id *string, options *ConditionalTriggerOptions) interfacesawsglue.ITriggerRef {
	if err := i.validateAddConditionalTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns interfacesawsglue.ITriggerRef

	_jsii_.Invoke(
		i,
		"addConditionalTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkflow) AddEventTrigger(id *string, options *EventTriggerOptions) interfacesawsglue.ITriggerRef {
	if err := i.validateAddEventTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns interfacesawsglue.ITriggerRef

	_jsii_.Invoke(
		i,
		"addEventTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkflow) AddOnDemandTrigger(id *string, options *OnDemandTriggerOptions) interfacesawsglue.ITriggerRef {
	if err := i.validateAddOnDemandTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns interfacesawsglue.ITriggerRef

	_jsii_.Invoke(
		i,
		"addOnDemandTrigger",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkflow) AddScheduledTrigger(id *string, options *ScheduledTriggerOptions) interfacesawsglue.ITriggerRef {
	if err := i.validateAddScheduledTriggerParameters(id, options); err != nil {
		panic(err)
	}
	var returns interfacesawsglue.ITriggerRef

	_jsii_.Invoke(
		i,
		"addScheduledTrigger",
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

