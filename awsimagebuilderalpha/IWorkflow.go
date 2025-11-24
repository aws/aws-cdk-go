package awsimagebuilderalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/internal"
)

// An EC2 Image Builder Workflow.
// Experimental.
type IWorkflow interface {
	awscdk.IResource
	// Grant custom actions to the given grantee for the workflow.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant read permissions to the given grantee for the workflow.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the workflow.
	// Experimental.
	WorkflowArn() *string
	// The name of the workflow.
	// Experimental.
	WorkflowName() *string
	// The type of the workflow.
	// Experimental.
	WorkflowType() *string
	// The version of the workflow.
	// Experimental.
	WorkflowVersion() *string
}

// The jsii proxy for IWorkflow
type jsiiProxy_IWorkflow struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IWorkflow) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkflow) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
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

func (j *jsiiProxy_IWorkflow) WorkflowType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkflow) WorkflowVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowVersion",
		&returns,
	)
	return returns
}

