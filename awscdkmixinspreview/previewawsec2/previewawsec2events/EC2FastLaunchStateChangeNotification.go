package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ec2@EC2FastLaunchStateChangeNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2FastLaunchStateChangeNotification := awscdkmixinspreview.Events.NewEC2FastLaunchStateChangeNotification()
//
// Experimental.
type EC2FastLaunchStateChangeNotification interface {
}

// The jsii proxy struct for EC2FastLaunchStateChangeNotification
type jsiiProxy_EC2FastLaunchStateChangeNotification struct {
	_ byte // padding
}

// Experimental.
func NewEC2FastLaunchStateChangeNotification() EC2FastLaunchStateChangeNotification {
	_init_.Initialize()

	j := jsiiProxy_EC2FastLaunchStateChangeNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EC2FastLaunchStateChangeNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2FastLaunchStateChangeNotification_Override(e EC2FastLaunchStateChangeNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EC2FastLaunchStateChangeNotification",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Fast Launch State-change Notification.
// Experimental.
func EC2FastLaunchStateChangeNotification_Ec2FastLaunchStateChangeNotificationPattern(options *EC2FastLaunchStateChangeNotification_EC2FastLaunchStateChangeNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2FastLaunchStateChangeNotification_Ec2FastLaunchStateChangeNotificationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.events.EC2FastLaunchStateChangeNotification",
		"ec2FastLaunchStateChangeNotificationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

