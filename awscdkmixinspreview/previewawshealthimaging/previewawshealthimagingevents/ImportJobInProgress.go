package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImportJobInProgress.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobInProgress := awscdkmixinspreview.Events.NewImportJobInProgress()
//
// Experimental.
type ImportJobInProgress interface {
}

// The jsii proxy struct for ImportJobInProgress
type jsiiProxy_ImportJobInProgress struct {
	_ byte // padding
}

// Experimental.
func NewImportJobInProgress() ImportJobInProgress {
	_init_.Initialize()

	j := jsiiProxy_ImportJobInProgress{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobInProgress",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImportJobInProgress_Override(i ImportJobInProgress) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobInProgress",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Import Job In Progress.
// Experimental.
func ImportJobInProgress_ImportJobInProgressPattern(options *ImportJobInProgress_ImportJobInProgressProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImportJobInProgress_ImportJobInProgressPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobInProgress",
		"importJobInProgressPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

