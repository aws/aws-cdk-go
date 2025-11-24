package mixinsawsdevopsguru

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnLogAnomalyDetectionIntegrationMixinProps",
		reflect.TypeOf((*CfnLogAnomalyDetectionIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnLogAnomalyDetectionIntegrationPropsMixin",
		reflect.TypeOf((*CfnLogAnomalyDetectionIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogAnomalyDetectionIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnNotificationChannelMixinProps",
		reflect.TypeOf((*CfnNotificationChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnNotificationChannelPropsMixin",
		reflect.TypeOf((*CfnNotificationChannelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNotificationChannelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnNotificationChannelPropsMixin.NotificationChannelConfigProperty",
		reflect.TypeOf((*CfnNotificationChannelPropsMixin_NotificationChannelConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnNotificationChannelPropsMixin.NotificationFilterConfigProperty",
		reflect.TypeOf((*CfnNotificationChannelPropsMixin_NotificationFilterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnNotificationChannelPropsMixin.SnsChannelConfigProperty",
		reflect.TypeOf((*CfnNotificationChannelPropsMixin_SnsChannelConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnResourceCollectionMixinProps",
		reflect.TypeOf((*CfnResourceCollectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnResourceCollectionPropsMixin",
		reflect.TypeOf((*CfnResourceCollectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourceCollectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnResourceCollectionPropsMixin.CloudFormationCollectionFilterProperty",
		reflect.TypeOf((*CfnResourceCollectionPropsMixin_CloudFormationCollectionFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnResourceCollectionPropsMixin.ResourceCollectionFilterProperty",
		reflect.TypeOf((*CfnResourceCollectionPropsMixin_ResourceCollectionFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsguru.mixins.CfnResourceCollectionPropsMixin.TagCollectionProperty",
		reflect.TypeOf((*CfnResourceCollectionPropsMixin_TagCollectionProperty)(nil)).Elem(),
	)
}
