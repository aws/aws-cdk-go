package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@TaskStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taskStatusChange := awscdkmixinspreview.Events.NewTaskStatusChange()
//
// Experimental.
type TaskStatusChange interface {
}

// The jsii proxy struct for TaskStatusChange
type jsiiProxy_TaskStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewTaskStatusChange() TaskStatusChange {
	_init_.Initialize()

	j := jsiiProxy_TaskStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.TaskStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTaskStatusChange_Override(t TaskStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.TaskStatusChange",
		nil, // no parameters
		t,
	)
}

// EventBridge event pattern for Task Status Change.
// Experimental.
func TaskStatusChange_TaskStatusChangePattern(options *TaskStatusChange_TaskStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateTaskStatusChange_TaskStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.TaskStatusChange",
		"taskStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

