package awsmgn

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnConnectorMixinProps",
		reflect.TypeOf((*CfnConnectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnConnectorPropsMixin",
		reflect.TypeOf((*CfnConnectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnConnectorPropsMixin.ConnectorSsmCommandConfigProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_ConnectorSsmCommandConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplateMixinProps",
		reflect.TypeOf((*CfnLaunchConfigurationTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin",
		reflect.TypeOf((*CfnLaunchConfigurationTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLaunchConfigurationTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin.LaunchTemplateDiskConfProperty",
		reflect.TypeOf((*CfnLaunchConfigurationTemplatePropsMixin_LaunchTemplateDiskConfProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin.LicensingProperty",
		reflect.TypeOf((*CfnLaunchConfigurationTemplatePropsMixin_LicensingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin.PostLaunchActionsProperty",
		reflect.TypeOf((*CfnLaunchConfigurationTemplatePropsMixin_PostLaunchActionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin.SsmDocumentProperty",
		reflect.TypeOf((*CfnLaunchConfigurationTemplatePropsMixin_SsmDocumentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin.SsmExternalParameterProperty",
		reflect.TypeOf((*CfnLaunchConfigurationTemplatePropsMixin_SsmExternalParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin.SsmParameterStoreParameterProperty",
		reflect.TypeOf((*CfnLaunchConfigurationTemplatePropsMixin_SsmParameterStoreParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnNetworkMigrationDefinitionMixinProps",
		reflect.TypeOf((*CfnNetworkMigrationDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnNetworkMigrationDefinitionPropsMixin",
		reflect.TypeOf((*CfnNetworkMigrationDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkMigrationDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnNetworkMigrationDefinitionPropsMixin.SourceConfigurationProperty",
		reflect.TypeOf((*CfnNetworkMigrationDefinitionPropsMixin_SourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnNetworkMigrationDefinitionPropsMixin.SourceS3ConfigurationProperty",
		reflect.TypeOf((*CfnNetworkMigrationDefinitionPropsMixin_SourceS3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnNetworkMigrationDefinitionPropsMixin.TargetNetworkProperty",
		reflect.TypeOf((*CfnNetworkMigrationDefinitionPropsMixin_TargetNetworkProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnNetworkMigrationDefinitionPropsMixin.TargetS3ConfigurationProperty",
		reflect.TypeOf((*CfnNetworkMigrationDefinitionPropsMixin_TargetS3ConfigurationProperty)(nil)).Elem(),
	)
}
