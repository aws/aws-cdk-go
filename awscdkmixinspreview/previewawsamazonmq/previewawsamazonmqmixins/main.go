package previewawsamazonmqmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerMixinProps",
		reflect.TypeOf((*CfnBrokerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin",
		reflect.TypeOf((*CfnBrokerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrokerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin.ConfigurationIdProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_ConfigurationIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin.EncryptionOptionsProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_EncryptionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin.LdapServerMetadataProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_LdapServerMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin.LogListProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_LogListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin.MaintenanceWindowProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_MaintenanceWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin.TagsEntryProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_TagsEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin.UserProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_UserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnConfigurationAssociationMixinProps",
		reflect.TypeOf((*CfnConfigurationAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnConfigurationAssociationPropsMixin",
		reflect.TypeOf((*CfnConfigurationAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnConfigurationAssociationPropsMixin.ConfigurationIdProperty",
		reflect.TypeOf((*CfnConfigurationAssociationPropsMixin_ConfigurationIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnConfigurationMixinProps",
		reflect.TypeOf((*CfnConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnConfigurationPropsMixin",
		reflect.TypeOf((*CfnConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnConfigurationPropsMixin.TagsEntryProperty",
		reflect.TypeOf((*CfnConfigurationPropsMixin_TagsEntryProperty)(nil)).Elem(),
	)
}
