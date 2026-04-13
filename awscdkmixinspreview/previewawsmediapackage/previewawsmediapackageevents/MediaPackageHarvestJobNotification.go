package previewawsmediapackageevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.mediapackage@MediaPackageHarvestJobNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaPackageHarvestJobNotification := awscdkmixinspreview.Events.NewMediaPackageHarvestJobNotification()
//
// Experimental.
type MediaPackageHarvestJobNotification interface {
}

// The jsii proxy struct for MediaPackageHarvestJobNotification
type jsiiProxy_MediaPackageHarvestJobNotification struct {
	_ byte // padding
}

// Experimental.
func NewMediaPackageHarvestJobNotification() MediaPackageHarvestJobNotification {
	_init_.Initialize()

	j := jsiiProxy_MediaPackageHarvestJobNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageHarvestJobNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMediaPackageHarvestJobNotification_Override(m MediaPackageHarvestJobNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageHarvestJobNotification",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for MediaPackage HarvestJob Notification.
// Experimental.
func MediaPackageHarvestJobNotification_EventPattern(options *MediaPackageHarvestJobNotification_MediaPackageHarvestJobNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateMediaPackageHarvestJobNotification_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageHarvestJobNotification",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

