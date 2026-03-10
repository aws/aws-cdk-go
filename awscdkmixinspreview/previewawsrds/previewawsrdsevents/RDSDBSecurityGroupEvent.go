package previewawsrdsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.rds@RDSDBSecurityGroupEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rDSDBSecurityGroupEvent := awscdkmixinspreview.Events.NewRDSDBSecurityGroupEvent()
//
// Experimental.
type RDSDBSecurityGroupEvent interface {
}

// The jsii proxy struct for RDSDBSecurityGroupEvent
type jsiiProxy_RDSDBSecurityGroupEvent struct {
	_ byte // padding
}

// Experimental.
func NewRDSDBSecurityGroupEvent() RDSDBSecurityGroupEvent {
	_init_.Initialize()

	j := jsiiProxy_RDSDBSecurityGroupEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBSecurityGroupEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRDSDBSecurityGroupEvent_Override(r RDSDBSecurityGroupEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBSecurityGroupEvent",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for RDS DB Security Group Event.
// Experimental.
func RDSDBSecurityGroupEvent_RdsDBSecurityGroupEventPattern(options *RDSDBSecurityGroupEvent_RDSDBSecurityGroupEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRDSDBSecurityGroupEvent_RdsDBSecurityGroupEventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.events.RDSDBSecurityGroupEvent",
		"rdsDBSecurityGroupEventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

