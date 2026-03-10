package previewawsmediapackageevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageHarvestJobNotification",
		reflect.TypeOf((*MediaPackageHarvestJobNotification)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MediaPackageHarvestJobNotification{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageHarvestJobNotification.HarvestJob",
		reflect.TypeOf((*MediaPackageHarvestJobNotification_HarvestJob)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageHarvestJobNotification.MediaPackageHarvestJobNotificationProps",
		reflect.TypeOf((*MediaPackageHarvestJobNotification_MediaPackageHarvestJobNotificationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageHarvestJobNotification.S3Destination",
		reflect.TypeOf((*MediaPackageHarvestJobNotification_S3Destination)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageInputNotification",
		reflect.TypeOf((*MediaPackageInputNotification)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MediaPackageInputNotification{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageInputNotification.MediaPackageInputNotificationProps",
		reflect.TypeOf((*MediaPackageInputNotification_MediaPackageInputNotificationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageKeyProviderNotification",
		reflect.TypeOf((*MediaPackageKeyProviderNotification)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MediaPackageKeyProviderNotification{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackage.events.MediaPackageKeyProviderNotification.MediaPackageKeyProviderNotificationProps",
		reflect.TypeOf((*MediaPackageKeyProviderNotification_MediaPackageKeyProviderNotificationProps)(nil)).Elem(),
	)
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
}
