package previewawshealthlakeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthlake@ImportJobCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobCompleted := awscdkmixinspreview.Events.NewImportJobCompleted()
//
// Experimental.
type ImportJobCompleted interface {
}

// The jsii proxy struct for ImportJobCompleted
type jsiiProxy_ImportJobCompleted struct {
	_ byte // padding
}

// Experimental.
func NewImportJobCompleted() ImportJobCompleted {
	_init_.Initialize()

	j := jsiiProxy_ImportJobCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ImportJobCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImportJobCompleted_Override(i ImportJobCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ImportJobCompleted",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Import Job Completed.
// Experimental.
func ImportJobCompleted_ImportJobCompletedPattern(options *ImportJobCompleted_ImportJobCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImportJobCompleted_ImportJobCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthlake.events.ImportJobCompleted",
		"importJobCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

