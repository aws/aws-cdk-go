package previewawsemrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.emr@EMRStepStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRStepStatusChange := awscdkmixinspreview.Events.NewEMRStepStatusChange()
//
// Experimental.
type EMRStepStatusChange interface {
}

// The jsii proxy struct for EMRStepStatusChange
type jsiiProxy_EMRStepStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewEMRStepStatusChange() EMRStepStatusChange {
	_init_.Initialize()

	j := jsiiProxy_EMRStepStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRStepStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEMRStepStatusChange_Override(e EMRStepStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRStepStatusChange",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EMR Step Status Change.
// Experimental.
func EMRStepStatusChange_EventPattern(options *EMRStepStatusChange_EMRStepStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEMRStepStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRStepStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

