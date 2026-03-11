package awsamazonmq

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnBrokerMixinProps",
		reflect.TypeOf((*CfnBrokerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnBrokerPropsMixin",
		reflect.TypeOf((*CfnBrokerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrokerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnBrokerPropsMixin.ConfigurationIdProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_ConfigurationIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnBrokerPropsMixin.EncryptionOptionsProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_EncryptionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnBrokerPropsMixin.LdapServerMetadataProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_LdapServerMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnBrokerPropsMixin.LogListProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_LogListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnBrokerPropsMixin.MaintenanceWindowProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_MaintenanceWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnBrokerPropsMixin.TagsEntryProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_TagsEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnBrokerPropsMixin.UserProperty",
		reflect.TypeOf((*CfnBrokerPropsMixin_UserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnConfigurationAssociationMixinProps",
		reflect.TypeOf((*CfnConfigurationAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnConfigurationAssociationPropsMixin",
		reflect.TypeOf((*CfnConfigurationAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnConfigurationAssociationPropsMixin.ConfigurationIdProperty",
		reflect.TypeOf((*CfnConfigurationAssociationPropsMixin_ConfigurationIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnConfigurationMixinProps",
		reflect.TypeOf((*CfnConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnConfigurationPropsMixin",
		reflect.TypeOf((*CfnConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amazonmq.CfnConfigurationPropsMixin.TagsEntryProperty",
		reflect.TypeOf((*CfnConfigurationPropsMixin_TagsEntryProperty)(nil)).Elem(),
	)
}
