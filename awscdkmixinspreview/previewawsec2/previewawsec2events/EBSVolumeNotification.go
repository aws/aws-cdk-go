package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ec2@EBSVolumeNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eBSVolumeNotification := awscdkmixinspreview.Events.NewEBSVolumeNotification()
//
// Experimental.
type EBSVolumeNotification interface {
}

// The jsii proxy struct for EBSVolumeNotification
type jsiiProxy_EBSVolumeNotification struct {
	_ byte // padding
}

// Experimental.
func NewEBSVolumeNotification() EBSVolumeNotification {
	_init_.Initialize()

	j := jsiiProxy_EBSVolumeNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EBSVolumeNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEBSVolumeNotification_Override(e EBSVolumeNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EBSVolumeNotification",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EBS Volume Notification.
// Experimental.
func EBSVolumeNotification_EbsVolumeNotificationPattern(options *EBSVolumeNotification_EBSVolumeNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEBSVolumeNotification_EbsVolumeNotificationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.events.EBSVolumeNotification",
		"ebsVolumeNotificationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

