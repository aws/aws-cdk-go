package previewawshealthlakeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthlake@ExportJobCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   exportJobCompleted := awscdkmixinspreview.Events.NewExportJobCompleted()
//
// Experimental.
type ExportJobCompleted interface {
}

// The jsii proxy struct for ExportJobCompleted
type jsiiProxy_ExportJobCompleted struct {
	_ byte // padding
}

// Experimental.
func NewExportJobCompleted() ExportJobCompleted {
	_init_.Initialize()

	j := jsiiProxy_ExportJobCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewExportJobCompleted_Override(e ExportJobCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobCompleted",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for Export Job Completed.
// Experimental.
func ExportJobCompleted_EventPattern(options *ExportJobCompleted_ExportJobCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateExportJobCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

