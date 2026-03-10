package previewawsemrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.emr@EMRConfigurationError.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRConfigurationError := awscdkmixinspreview.Events.NewEMRConfigurationError()
//
// Experimental.
type EMRConfigurationError interface {
}

// The jsii proxy struct for EMRConfigurationError
type jsiiProxy_EMRConfigurationError struct {
	_ byte // padding
}

// Experimental.
func NewEMRConfigurationError() EMRConfigurationError {
	_init_.Initialize()

	j := jsiiProxy_EMRConfigurationError{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRConfigurationError",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEMRConfigurationError_Override(e EMRConfigurationError) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRConfigurationError",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EMR Configuration Error.
// Experimental.
func EMRConfigurationError_EmrConfigurationErrorPattern(options *EMRConfigurationError_EMRConfigurationErrorProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEMRConfigurationError_EmrConfigurationErrorPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRConfigurationError",
		"emrConfigurationErrorPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

