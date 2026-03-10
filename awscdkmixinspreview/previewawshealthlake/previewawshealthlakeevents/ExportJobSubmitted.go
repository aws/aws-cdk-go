package previewawshealthlakeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthlake@ExportJobSubmitted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   exportJobSubmitted := awscdkmixinspreview.Events.NewExportJobSubmitted()
//
// Experimental.
type ExportJobSubmitted interface {
}

// The jsii proxy struct for ExportJobSubmitted
type jsiiProxy_ExportJobSubmitted struct {
	_ byte // padding
}

// Experimental.
func NewExportJobSubmitted() ExportJobSubmitted {
	_init_.Initialize()

	j := jsiiProxy_ExportJobSubmitted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobSubmitted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewExportJobSubmitted_Override(e ExportJobSubmitted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobSubmitted",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for Export Job Submitted.
// Experimental.
func ExportJobSubmitted_ExportJobSubmittedPattern(options *ExportJobSubmitted_ExportJobSubmittedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateExportJobSubmitted_ExportJobSubmittedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ExportJobSubmitted",
		"exportJobSubmittedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

