package previewawsmediapackageevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.mediapackage@MediaPackageKeyProviderNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaPackageKeyProviderNotification := awscdkmixinspreview.Events.NewMediaPackageKeyProviderNotification()
//
// Experimental.
type MediaPackageKeyProviderNotification interface {
}

// The jsii proxy struct for MediaPackageKeyProviderNotification
type jsiiProxy_MediaPackageKeyProviderNotification struct {
	_ byte // padding
}

// Experimental.
func NewMediaPackageKeyProviderNotification() MediaPackageKeyProviderNotification {
	_init_.Initialize()

	j := jsiiProxy_MediaPackageKeyProviderNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageKeyProviderNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMediaPackageKeyProviderNotification_Override(m MediaPackageKeyProviderNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageKeyProviderNotification",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MediaPackage Key Provider Notification.
// Experimental.
func MediaPackageKeyProviderNotification_MediaPackageKeyProviderNotificationPattern(options *MediaPackageKeyProviderNotification_MediaPackageKeyProviderNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMediaPackageKeyProviderNotification_MediaPackageKeyProviderNotificationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageKeyProviderNotification",
		"mediaPackageKeyProviderNotificationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

