package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@WorkflowStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workflowStatusChange := awscdkmixinspreview.Events.NewWorkflowStatusChange()
//
// Experimental.
type WorkflowStatusChange interface {
}

// The jsii proxy struct for WorkflowStatusChange
type jsiiProxy_WorkflowStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewWorkflowStatusChange() WorkflowStatusChange {
	_init_.Initialize()

	j := jsiiProxy_WorkflowStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.WorkflowStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWorkflowStatusChange_Override(w WorkflowStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.WorkflowStatusChange",
		nil, // no parameters
		w,
	)
}

// EventBridge event pattern for Workflow Status Change.
// Experimental.
func WorkflowStatusChange_WorkflowStatusChangePattern(options *WorkflowStatusChange_WorkflowStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateWorkflowStatusChange_WorkflowStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.WorkflowStatusChange",
		"workflowStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

