package awssecuritylake

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnAwsLogSourceMixinProps",
		reflect.TypeOf((*CfnAwsLogSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnAwsLogSourcePropsMixin",
		reflect.TypeOf((*CfnAwsLogSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAwsLogSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnDataLakeMixinProps",
		reflect.TypeOf((*CfnDataLakeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnDataLakePropsMixin",
		reflect.TypeOf((*CfnDataLakePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataLakePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnDataLakePropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnDataLakePropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnDataLakePropsMixin.ExpirationProperty",
		reflect.TypeOf((*CfnDataLakePropsMixin_ExpirationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnDataLakePropsMixin.LifecycleConfigurationProperty",
		reflect.TypeOf((*CfnDataLakePropsMixin_LifecycleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnDataLakePropsMixin.ReplicationConfigurationProperty",
		reflect.TypeOf((*CfnDataLakePropsMixin_ReplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnDataLakePropsMixin.TransitionsProperty",
		reflect.TypeOf((*CfnDataLakePropsMixin_TransitionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnSubscriberMixinProps",
		reflect.TypeOf((*CfnSubscriberMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnSubscriberNotificationMixinProps",
		reflect.TypeOf((*CfnSubscriberNotificationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnSubscriberNotificationPropsMixin",
		reflect.TypeOf((*CfnSubscriberNotificationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubscriberNotificationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnSubscriberNotificationPropsMixin.HttpsNotificationConfigurationProperty",
		reflect.TypeOf((*CfnSubscriberNotificationPropsMixin_HttpsNotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnSubscriberNotificationPropsMixin.NotificationConfigurationProperty",
		reflect.TypeOf((*CfnSubscriberNotificationPropsMixin_NotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnSubscriberPropsMixin",
		reflect.TypeOf((*CfnSubscriberPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubscriberPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnSubscriberPropsMixin.AwsLogSourceProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_AwsLogSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnSubscriberPropsMixin.CustomLogSourceProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_CustomLogSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnSubscriberPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securitylake.CfnSubscriberPropsMixin.SubscriberIdentityProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_SubscriberIdentityProperty)(nil)).Elem(),
	)
}
