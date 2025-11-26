package previewawsmediapackageevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackage"
)

// EventBridge event patterns for OriginEndpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var originEndpointRef IOriginEndpointRef
//
//   originEndpointEvents := awscdkmixinspreview.Events.OriginEndpointEvents_FromOriginEndpoint(originEndpointRef)
//
// Experimental.
type OriginEndpointEvents interface {
	// EventBridge event pattern for OriginEndpoint MediaPackage HarvestJob Notification.
	// Experimental.
	MediaPackageHarvestJobNotificationPattern(options *OriginEndpointEvents_MediaPackageHarvestJobNotification_MediaPackageHarvestJobNotificationProps) *awsevents.EventPattern
}

// The jsii proxy struct for OriginEndpointEvents
type jsiiProxy_OriginEndpointEvents struct {
	_ byte // padding
}

// Create OriginEndpointEvents from a OriginEndpoint reference.
// Experimental.
func OriginEndpointEvents_FromOriginEndpoint(originEndpointRef interfacesawsmediapackage.IOriginEndpointRef) OriginEndpointEvents {
	_init_.Initialize()

	if err := validateOriginEndpointEvents_FromOriginEndpointParameters(originEndpointRef); err != nil {
		panic(err)
	}
	var returns OriginEndpointEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.OriginEndpointEvents",
		"fromOriginEndpoint",
		[]interface{}{originEndpointRef},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginEndpointEvents) MediaPackageHarvestJobNotificationPattern(options *OriginEndpointEvents_MediaPackageHarvestJobNotification_MediaPackageHarvestJobNotificationProps) *awsevents.EventPattern {
	if err := o.validateMediaPackageHarvestJobNotificationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		o,
		"mediaPackageHarvestJobNotificationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

