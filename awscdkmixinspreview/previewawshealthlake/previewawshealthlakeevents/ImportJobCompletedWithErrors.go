package previewawshealthlakeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthlake@ImportJobCompletedWithErrors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobCompletedWithErrors := awscdkmixinspreview.Events.NewImportJobCompletedWithErrors()
//
// Experimental.
type ImportJobCompletedWithErrors interface {
}

// The jsii proxy struct for ImportJobCompletedWithErrors
type jsiiProxy_ImportJobCompletedWithErrors struct {
	_ byte // padding
}

// Experimental.
func NewImportJobCompletedWithErrors() ImportJobCompletedWithErrors {
	_init_.Initialize()

	j := jsiiProxy_ImportJobCompletedWithErrors{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ImportJobCompletedWithErrors",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImportJobCompletedWithErrors_Override(i ImportJobCompletedWithErrors) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ImportJobCompletedWithErrors",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Import Job Completed With Errors.
// Experimental.
func ImportJobCompletedWithErrors_EventPattern(options *ImportJobCompletedWithErrors_ImportJobCompletedWithErrorsProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImportJobCompletedWithErrors_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ImportJobCompletedWithErrors",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

