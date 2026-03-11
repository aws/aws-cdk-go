package awsdevopsguru

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnLogAnomalyDetectionIntegrationMixinProps",
		reflect.TypeOf((*CfnLogAnomalyDetectionIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnLogAnomalyDetectionIntegrationPropsMixin",
		reflect.TypeOf((*CfnLogAnomalyDetectionIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogAnomalyDetectionIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnNotificationChannelMixinProps",
		reflect.TypeOf((*CfnNotificationChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnNotificationChannelPropsMixin",
		reflect.TypeOf((*CfnNotificationChannelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNotificationChannelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnNotificationChannelPropsMixin.NotificationChannelConfigProperty",
		reflect.TypeOf((*CfnNotificationChannelPropsMixin_NotificationChannelConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnNotificationChannelPropsMixin.NotificationFilterConfigProperty",
		reflect.TypeOf((*CfnNotificationChannelPropsMixin_NotificationFilterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnNotificationChannelPropsMixin.SnsChannelConfigProperty",
		reflect.TypeOf((*CfnNotificationChannelPropsMixin_SnsChannelConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnResourceCollectionMixinProps",
		reflect.TypeOf((*CfnResourceCollectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnResourceCollectionPropsMixin",
		reflect.TypeOf((*CfnResourceCollectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourceCollectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnResourceCollectionPropsMixin.CloudFormationCollectionFilterProperty",
		reflect.TypeOf((*CfnResourceCollectionPropsMixin_CloudFormationCollectionFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnResourceCollectionPropsMixin.ResourceCollectionFilterProperty",
		reflect.TypeOf((*CfnResourceCollectionPropsMixin_ResourceCollectionFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsguru.CfnResourceCollectionPropsMixin.TagCollectionProperty",
		reflect.TypeOf((*CfnResourceCollectionPropsMixin_TagCollectionProperty)(nil)).Elem(),
	)
}
