package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImportJobFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobFailed := awscdkmixinspreview.Events.NewImportJobFailed()
//
// Experimental.
type ImportJobFailed interface {
}

// The jsii proxy struct for ImportJobFailed
type jsiiProxy_ImportJobFailed struct {
	_ byte // padding
}

// Experimental.
func NewImportJobFailed() ImportJobFailed {
	_init_.Initialize()

	j := jsiiProxy_ImportJobFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImportJobFailed_Override(i ImportJobFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobFailed",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Import Job Failed.
// Experimental.
func ImportJobFailed_ImportJobFailedPattern(options *ImportJobFailed_ImportJobFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImportJobFailed_ImportJobFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobFailed",
		"importJobFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

