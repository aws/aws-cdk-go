package previewawsemrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.emr@EMRInstanceFleetStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRInstanceFleetStateChange := awscdkmixinspreview.Events.NewEMRInstanceFleetStateChange()
//
// Experimental.
type EMRInstanceFleetStateChange interface {
}

// The jsii proxy struct for EMRInstanceFleetStateChange
type jsiiProxy_EMRInstanceFleetStateChange struct {
	_ byte // padding
}

// Experimental.
func NewEMRInstanceFleetStateChange() EMRInstanceFleetStateChange {
	_init_.Initialize()

	j := jsiiProxy_EMRInstanceFleetStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRInstanceFleetStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEMRInstanceFleetStateChange_Override(e EMRInstanceFleetStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRInstanceFleetStateChange",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EMR Instance Fleet State Change.
// Experimental.
func EMRInstanceFleetStateChange_EmrInstanceFleetStateChangePattern(options *EMRInstanceFleetStateChange_EMRInstanceFleetStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEMRInstanceFleetStateChange_EmrInstanceFleetStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRInstanceFleetStateChange",
		"emrInstanceFleetStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

