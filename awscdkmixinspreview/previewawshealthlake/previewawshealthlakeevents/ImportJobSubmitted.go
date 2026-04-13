package previewawshealthlakeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthlake@ImportJobSubmitted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobSubmitted := awscdkmixinspreview.Events.NewImportJobSubmitted()
//
// Experimental.
type ImportJobSubmitted interface {
}

// The jsii proxy struct for ImportJobSubmitted
type jsiiProxy_ImportJobSubmitted struct {
	_ byte // padding
}

// Experimental.
func NewImportJobSubmitted() ImportJobSubmitted {
	_init_.Initialize()

	j := jsiiProxy_ImportJobSubmitted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ImportJobSubmitted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImportJobSubmitted_Override(i ImportJobSubmitted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ImportJobSubmitted",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Import Job Submitted.
// Experimental.
func ImportJobSubmitted_EventPattern(options *ImportJobSubmitted_ImportJobSubmittedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImportJobSubmitted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ImportJobSubmitted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

