package previewawsemrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.emr@EMRInstanceGroupStatusNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRInstanceGroupStatusNotification := awscdkmixinspreview.Events.NewEMRInstanceGroupStatusNotification()
//
// Experimental.
type EMRInstanceGroupStatusNotification interface {
}

// The jsii proxy struct for EMRInstanceGroupStatusNotification
type jsiiProxy_EMRInstanceGroupStatusNotification struct {
	_ byte // padding
}

// Experimental.
func NewEMRInstanceGroupStatusNotification() EMRInstanceGroupStatusNotification {
	_init_.Initialize()

	j := jsiiProxy_EMRInstanceGroupStatusNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRInstanceGroupStatusNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEMRInstanceGroupStatusNotification_Override(e EMRInstanceGroupStatusNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRInstanceGroupStatusNotification",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EMR Instance Group Status Notification.
// Experimental.
func EMRInstanceGroupStatusNotification_EventPattern(options *EMRInstanceGroupStatusNotification_EMRInstanceGroupStatusNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEMRInstanceGroupStatusNotification_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRInstanceGroupStatusNotification",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

