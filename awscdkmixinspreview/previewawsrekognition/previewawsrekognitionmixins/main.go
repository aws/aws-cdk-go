package previewawsrekognitionmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnCollectionMixinProps",
		reflect.TypeOf((*CfnCollectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnCollectionPropsMixin",
		reflect.TypeOf((*CfnCollectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCollectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorMixinProps",
		reflect.TypeOf((*CfnStreamProcessorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStreamProcessorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin.BoundingBoxProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_BoundingBoxProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin.ConnectedHomeSettingsProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_ConnectedHomeSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin.DataSharingPreferenceProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_DataSharingPreferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin.FaceSearchSettingsProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_FaceSearchSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin.KinesisDataStreamProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_KinesisDataStreamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin.KinesisVideoStreamProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_KinesisVideoStreamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin.NotificationChannelProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_NotificationChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin.PointProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_PointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rekognition.mixins.CfnStreamProcessorPropsMixin.S3DestinationProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_S3DestinationProperty)(nil)).Elem(),
	)
}
