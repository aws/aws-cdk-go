package previewawsrdsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.rds@RDSDBParameterGroupEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rDSDBParameterGroupEvent := awscdkmixinspreview.Events.NewRDSDBParameterGroupEvent()
//
// Experimental.
type RDSDBParameterGroupEvent interface {
}

// The jsii proxy struct for RDSDBParameterGroupEvent
type jsiiProxy_RDSDBParameterGroupEvent struct {
	_ byte // padding
}

// Experimental.
func NewRDSDBParameterGroupEvent() RDSDBParameterGroupEvent {
	_init_.Initialize()

	j := jsiiProxy_RDSDBParameterGroupEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBParameterGroupEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRDSDBParameterGroupEvent_Override(r RDSDBParameterGroupEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBParameterGroupEvent",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for RDS DB Parameter Group Event.
// Experimental.
func RDSDBParameterGroupEvent_EventPattern(options *RDSDBParameterGroupEvent_RDSDBParameterGroupEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRDSDBParameterGroupEvent_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBParameterGroupEvent",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

