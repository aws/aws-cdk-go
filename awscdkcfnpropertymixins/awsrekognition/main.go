package awsrekognition

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnCollectionMixinProps",
		reflect.TypeOf((*CfnCollectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnCollectionPropsMixin",
		reflect.TypeOf((*CfnCollectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCollectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorMixinProps",
		reflect.TypeOf((*CfnStreamProcessorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorPropsMixin",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStreamProcessorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorPropsMixin.BoundingBoxProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_BoundingBoxProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorPropsMixin.ConnectedHomeSettingsProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_ConnectedHomeSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorPropsMixin.DataSharingPreferenceProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_DataSharingPreferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorPropsMixin.FaceSearchSettingsProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_FaceSearchSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorPropsMixin.KinesisDataStreamProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_KinesisDataStreamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorPropsMixin.KinesisVideoStreamProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_KinesisVideoStreamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorPropsMixin.NotificationChannelProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_NotificationChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorPropsMixin.PointProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_PointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rekognition.CfnStreamProcessorPropsMixin.S3DestinationProperty",
		reflect.TypeOf((*CfnStreamProcessorPropsMixin_S3DestinationProperty)(nil)).Elem(),
	)
}
