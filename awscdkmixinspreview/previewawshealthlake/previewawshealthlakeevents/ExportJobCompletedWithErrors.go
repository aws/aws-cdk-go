package previewawshealthlakeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthlake@ExportJobCompletedWithErrors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   exportJobCompletedWithErrors := awscdkmixinspreview.Events.NewExportJobCompletedWithErrors()
//
// Experimental.
type ExportJobCompletedWithErrors interface {
}

// The jsii proxy struct for ExportJobCompletedWithErrors
type jsiiProxy_ExportJobCompletedWithErrors struct {
	_ byte // padding
}

// Experimental.
func NewExportJobCompletedWithErrors() ExportJobCompletedWithErrors {
	_init_.Initialize()

	j := jsiiProxy_ExportJobCompletedWithErrors{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobCompletedWithErrors",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewExportJobCompletedWithErrors_Override(e ExportJobCompletedWithErrors) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobCompletedWithErrors",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for Export Job Completed With Errors.
// Experimental.
func ExportJobCompletedWithErrors_EventPattern(options *ExportJobCompletedWithErrors_ExportJobCompletedWithErrorsProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateExportJobCompletedWithErrors_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobCompletedWithErrors",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

