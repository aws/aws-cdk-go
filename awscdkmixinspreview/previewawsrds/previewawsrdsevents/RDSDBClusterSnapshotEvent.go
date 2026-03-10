package previewawsrdsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.rds@RDSDBClusterSnapshotEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rDSDBClusterSnapshotEvent := awscdkmixinspreview.Events.NewRDSDBClusterSnapshotEvent()
//
// Experimental.
type RDSDBClusterSnapshotEvent interface {
}

// The jsii proxy struct for RDSDBClusterSnapshotEvent
type jsiiProxy_RDSDBClusterSnapshotEvent struct {
	_ byte // padding
}

// Experimental.
func NewRDSDBClusterSnapshotEvent() RDSDBClusterSnapshotEvent {
	_init_.Initialize()

	j := jsiiProxy_RDSDBClusterSnapshotEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBClusterSnapshotEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRDSDBClusterSnapshotEvent_Override(r RDSDBClusterSnapshotEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBClusterSnapshotEvent",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for RDS DB Cluster Snapshot Event.
// Experimental.
func RDSDBClusterSnapshotEvent_RdsDBClusterSnapshotEventPattern(options *RDSDBClusterSnapshotEvent_RDSDBClusterSnapshotEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRDSDBClusterSnapshotEvent_RdsDBClusterSnapshotEventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBClusterSnapshotEvent",
		"rdsDBClusterSnapshotEventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

