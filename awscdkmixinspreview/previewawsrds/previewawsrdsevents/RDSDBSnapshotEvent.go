package previewawsrdsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.rds@RDSDBSnapshotEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rDSDBSnapshotEvent := awscdkmixinspreview.Events.NewRDSDBSnapshotEvent()
//
// Experimental.
type RDSDBSnapshotEvent interface {
}

// The jsii proxy struct for RDSDBSnapshotEvent
type jsiiProxy_RDSDBSnapshotEvent struct {
	_ byte // padding
}

// Experimental.
func NewRDSDBSnapshotEvent() RDSDBSnapshotEvent {
	_init_.Initialize()

	j := jsiiProxy_RDSDBSnapshotEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBSnapshotEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRDSDBSnapshotEvent_Override(r RDSDBSnapshotEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBSnapshotEvent",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for RDS DB Snapshot Event.
// Experimental.
func RDSDBSnapshotEvent_EventPattern(options *RDSDBSnapshotEvent_RDSDBSnapshotEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRDSDBSnapshotEvent_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBSnapshotEvent",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

