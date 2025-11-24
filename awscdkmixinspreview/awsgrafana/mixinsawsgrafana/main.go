package mixinsawsgrafana

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspaceMixinProps",
		reflect.TypeOf((*CfnWorkspaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin",
		reflect.TypeOf((*CfnWorkspacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkspacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin.AssertionAttributesProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_AssertionAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin.IdpMetadataProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_IdpMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin.NetworkAccessControlProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_NetworkAccessControlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin.RoleValuesProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_RoleValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin.SamlConfigurationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_SamlConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_grafana.mixins.CfnWorkspacePropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
}
