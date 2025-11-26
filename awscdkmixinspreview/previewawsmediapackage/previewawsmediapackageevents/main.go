package previewawsmediapackageevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.OriginEndpointEvents",
		reflect.TypeOf((*OriginEndpointEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "mediaPackageHarvestJobNotificationPattern", GoMethod: "MediaPackageHarvestJobNotificationPattern"},
		},
		func() interface{} {
			return &jsiiProxy_OriginEndpointEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.OriginEndpointEvents.MediaPackageHarvestJobNotification",
		reflect.TypeOf((*OriginEndpointEvents_MediaPackageHarvestJobNotification)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_OriginEndpointEvents_MediaPackageHarvestJobNotification{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.OriginEndpointEvents.MediaPackageHarvestJobNotification.HarvestJob",
		reflect.TypeOf((*OriginEndpointEvents_MediaPackageHarvestJobNotification_HarvestJob)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.OriginEndpointEvents.MediaPackageHarvestJobNotification.MediaPackageHarvestJobNotificationProps",
		reflect.TypeOf((*OriginEndpointEvents_MediaPackageHarvestJobNotification_MediaPackageHarvestJobNotificationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.OriginEndpointEvents.MediaPackageHarvestJobNotification.S3Destination",
		reflect.TypeOf((*OriginEndpointEvents_MediaPackageHarvestJobNotification_S3Destination)(nil)).Elem(),
	)
}
