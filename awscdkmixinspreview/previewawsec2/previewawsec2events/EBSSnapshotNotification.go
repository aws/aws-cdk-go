package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ec2@EBSSnapshotNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eBSSnapshotNotification := awscdkmixinspreview.Events.NewEBSSnapshotNotification()
//
// Experimental.
type EBSSnapshotNotification interface {
}

// The jsii proxy struct for EBSSnapshotNotification
type jsiiProxy_EBSSnapshotNotification struct {
	_ byte // padding
}

// Experimental.
func NewEBSSnapshotNotification() EBSSnapshotNotification {
	_init_.Initialize()

	j := jsiiProxy_EBSSnapshotNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EBSSnapshotNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEBSSnapshotNotification_Override(e EBSSnapshotNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EBSSnapshotNotification",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EBS Snapshot Notification.
// Experimental.
func EBSSnapshotNotification_EventPattern(options *EBSSnapshotNotification_EBSSnapshotNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEBSSnapshotNotification_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.events.EBSSnapshotNotification",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

