package awsmgn

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
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
