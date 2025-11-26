package previewawsmediapackageevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.mediapackage@MediaPackageHarvestJobNotification event types for OriginEndpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaPackageHarvestJobNotification := #error#.NewMediaPackageHarvestJobNotification()
//
// Experimental.
type OriginEndpointEvents_MediaPackageHarvestJobNotification interface {
}

// The jsii proxy struct for OriginEndpointEvents_MediaPackageHarvestJobNotification
type jsiiProxy_OriginEndpointEvents_MediaPackageHarvestJobNotification struct {
	_ byte // padding
}

// Experimental.
func NewOriginEndpointEvents_MediaPackageHarvestJobNotification() OriginEndpointEvents_MediaPackageHarvestJobNotification {
	_init_.Initialize()

	j := jsiiProxy_OriginEndpointEvents_MediaPackageHarvestJobNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.OriginEndpointEvents.MediaPackageHarvestJobNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewOriginEndpointEvents_MediaPackageHarvestJobNotification_Override(o OriginEndpointEvents_MediaPackageHarvestJobNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.OriginEndpointEvents.MediaPackageHarvestJobNotification",
		nil, // no parameters
		o,
	)
}

