package previewawsrdsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.rds@RDSDBClusterEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rDSDBClusterEvent := awscdkmixinspreview.Events.NewRDSDBClusterEvent()
//
// Experimental.
type RDSDBClusterEvent interface {
}

// The jsii proxy struct for RDSDBClusterEvent
type jsiiProxy_RDSDBClusterEvent struct {
	_ byte // padding
}

// Experimental.
func NewRDSDBClusterEvent() RDSDBClusterEvent {
	_init_.Initialize()

	j := jsiiProxy_RDSDBClusterEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBClusterEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRDSDBClusterEvent_Override(r RDSDBClusterEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBClusterEvent",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for RDS DB Cluster Event.
// Experimental.
func RDSDBClusterEvent_RdsDBClusterEventPattern(options *RDSDBClusterEvent_RDSDBClusterEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRDSDBClusterEvent_RdsDBClusterEventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBClusterEvent",
		"rdsDBClusterEventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

