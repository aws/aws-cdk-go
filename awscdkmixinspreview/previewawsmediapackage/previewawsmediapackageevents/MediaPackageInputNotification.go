package previewawsmediapackageevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.mediapackage@MediaPackageInputNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaPackageInputNotification := awscdkmixinspreview.Events.NewMediaPackageInputNotification()
//
// Experimental.
type MediaPackageInputNotification interface {
}

// The jsii proxy struct for MediaPackageInputNotification
type jsiiProxy_MediaPackageInputNotification struct {
	_ byte // padding
}

// Experimental.
func NewMediaPackageInputNotification() MediaPackageInputNotification {
	_init_.Initialize()

	j := jsiiProxy_MediaPackageInputNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageInputNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMediaPackageInputNotification_Override(m MediaPackageInputNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageInputNotification",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MediaPackage Input Notification.
// Experimental.
func MediaPackageInputNotification_MediaPackageInputNotificationPattern(options *MediaPackageInputNotification_MediaPackageInputNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMediaPackageInputNotification_MediaPackageInputNotificationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageInputNotification",
		"mediaPackageInputNotificationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

