package previewawshealthlakeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthlake@ExportJobInProgress.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   exportJobInProgress := awscdkmixinspreview.Events.NewExportJobInProgress()
//
// Experimental.
type ExportJobInProgress interface {
}

// The jsii proxy struct for ExportJobInProgress
type jsiiProxy_ExportJobInProgress struct {
	_ byte // padding
}

// Experimental.
func NewExportJobInProgress() ExportJobInProgress {
	_init_.Initialize()

	j := jsiiProxy_ExportJobInProgress{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobInProgress",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewExportJobInProgress_Override(e ExportJobInProgress) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobInProgress",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for Export Job In Progress.
// Experimental.
func ExportJobInProgress_EventPattern(options *ExportJobInProgress_ExportJobInProgressProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateExportJobInProgress_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobInProgress",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

