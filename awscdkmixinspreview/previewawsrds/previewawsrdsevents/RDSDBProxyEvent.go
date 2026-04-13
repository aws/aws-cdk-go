package previewawsrdsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.rds@RDSDBProxyEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rDSDBProxyEvent := awscdkmixinspreview.Events.NewRDSDBProxyEvent()
//
// Experimental.
type RDSDBProxyEvent interface {
}

// The jsii proxy struct for RDSDBProxyEvent
type jsiiProxy_RDSDBProxyEvent struct {
	_ byte // padding
}

// Experimental.
func NewRDSDBProxyEvent() RDSDBProxyEvent {
	_init_.Initialize()

	j := jsiiProxy_RDSDBProxyEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBProxyEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRDSDBProxyEvent_Override(r RDSDBProxyEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBProxyEvent",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for RDS DB Proxy Event.
// Experimental.
func RDSDBProxyEvent_EventPattern(options *RDSDBProxyEvent_RDSDBProxyEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRDSDBProxyEvent_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBProxyEvent",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

