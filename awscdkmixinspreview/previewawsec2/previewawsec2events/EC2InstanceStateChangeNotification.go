package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ec2@EC2InstanceStateChangeNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceStateChangeNotification := awscdkmixinspreview.Events.NewEC2InstanceStateChangeNotification()
//
// Experimental.
type EC2InstanceStateChangeNotification interface {
}

// The jsii proxy struct for EC2InstanceStateChangeNotification
type jsiiProxy_EC2InstanceStateChangeNotification struct {
	_ byte // padding
}

// Experimental.
func NewEC2InstanceStateChangeNotification() EC2InstanceStateChangeNotification {
	_init_.Initialize()

	j := jsiiProxy_EC2InstanceStateChangeNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EC2InstanceStateChangeNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2InstanceStateChangeNotification_Override(e EC2InstanceStateChangeNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EC2InstanceStateChangeNotification",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Instance State-change Notification.
// Experimental.
func EC2InstanceStateChangeNotification_EventPattern(options *EC2InstanceStateChangeNotification_EC2InstanceStateChangeNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2InstanceStateChangeNotification_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.events.EC2InstanceStateChangeNotification",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

