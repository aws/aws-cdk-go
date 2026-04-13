package previewawsrdsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.rds@RDSDBInstanceEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rDSDBInstanceEvent := awscdkmixinspreview.Events.NewRDSDBInstanceEvent()
//
// Experimental.
type RDSDBInstanceEvent interface {
}

// The jsii proxy struct for RDSDBInstanceEvent
type jsiiProxy_RDSDBInstanceEvent struct {
	_ byte // padding
}

// Experimental.
func NewRDSDBInstanceEvent() RDSDBInstanceEvent {
	_init_.Initialize()

	j := jsiiProxy_RDSDBInstanceEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBInstanceEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRDSDBInstanceEvent_Override(r RDSDBInstanceEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBInstanceEvent",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for RDS DB Instance Event.
// Experimental.
func RDSDBInstanceEvent_EventPattern(options *RDSDBInstanceEvent_RDSDBInstanceEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRDSDBInstanceEvent_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBInstanceEvent",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

