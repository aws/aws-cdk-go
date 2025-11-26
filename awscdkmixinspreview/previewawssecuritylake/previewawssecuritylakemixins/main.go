package previewawssecuritylakemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnAwsLogSourceMixinProps",
		reflect.TypeOf((*CfnAwsLogSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnAwsLogSourcePropsMixin",
		reflect.TypeOf((*CfnAwsLogSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAwsLogSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakeMixinProps",
		reflect.TypeOf((*CfnDataLakeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakePropsMixin",
		reflect.TypeOf((*CfnDataLakePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataLakePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakePropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnDataLakePropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakePropsMixin.ExpirationProperty",
		reflect.TypeOf((*CfnDataLakePropsMixin_ExpirationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakePropsMixin.LifecycleConfigurationProperty",
		reflect.TypeOf((*CfnDataLakePropsMixin_LifecycleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakePropsMixin.ReplicationConfigurationProperty",
		reflect.TypeOf((*CfnDataLakePropsMixin_ReplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnDataLakePropsMixin.TransitionsProperty",
		reflect.TypeOf((*CfnDataLakePropsMixin_TransitionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberMixinProps",
		reflect.TypeOf((*CfnSubscriberMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberNotificationMixinProps",
		reflect.TypeOf((*CfnSubscriberNotificationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberNotificationPropsMixin",
		reflect.TypeOf((*CfnSubscriberNotificationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubscriberNotificationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberNotificationPropsMixin.HttpsNotificationConfigurationProperty",
		reflect.TypeOf((*CfnSubscriberNotificationPropsMixin_HttpsNotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberNotificationPropsMixin.NotificationConfigurationProperty",
		reflect.TypeOf((*CfnSubscriberNotificationPropsMixin_NotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberPropsMixin",
		reflect.TypeOf((*CfnSubscriberPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubscriberPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberPropsMixin.AwsLogSourceProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_AwsLogSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberPropsMixin.CustomLogSourceProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_CustomLogSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberPropsMixin.SubscriberIdentityProperty",
		reflect.TypeOf((*CfnSubscriberPropsMixin_SubscriberIdentityProperty)(nil)).Elem(),
	)
}
