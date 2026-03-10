package previewawshealthlakeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthlake@ExportJobFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   exportJobFailed := awscdkmixinspreview.Events.NewExportJobFailed()
//
// Experimental.
type ExportJobFailed interface {
}

// The jsii proxy struct for ExportJobFailed
type jsiiProxy_ExportJobFailed struct {
	_ byte // padding
}

// Experimental.
func NewExportJobFailed() ExportJobFailed {
	_init_.Initialize()

	j := jsiiProxy_ExportJobFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewExportJobFailed_Override(e ExportJobFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobFailed",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for Export Job Failed.
// Experimental.
func ExportJobFailed_ExportJobFailedPattern(options *ExportJobFailed_ExportJobFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateExportJobFailed_ExportJobFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobFailed",
		"exportJobFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

