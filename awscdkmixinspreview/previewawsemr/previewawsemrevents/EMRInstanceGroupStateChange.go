package previewawsemrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.emr@EMRInstanceGroupStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRInstanceGroupStateChange := awscdkmixinspreview.Events.NewEMRInstanceGroupStateChange()
//
// Experimental.
type EMRInstanceGroupStateChange interface {
}

// The jsii proxy struct for EMRInstanceGroupStateChange
type jsiiProxy_EMRInstanceGroupStateChange struct {
	_ byte // padding
}

// Experimental.
func NewEMRInstanceGroupStateChange() EMRInstanceGroupStateChange {
	_init_.Initialize()

	j := jsiiProxy_EMRInstanceGroupStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRInstanceGroupStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEMRInstanceGroupStateChange_Override(e EMRInstanceGroupStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRInstanceGroupStateChange",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EMR Instance Group State Change.
// Experimental.
func EMRInstanceGroupStateChange_EventPattern(options *EMRInstanceGroupStateChange_EMRInstanceGroupStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEMRInstanceGroupStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRInstanceGroupStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

