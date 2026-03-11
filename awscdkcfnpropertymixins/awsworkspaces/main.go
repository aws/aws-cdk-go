package awsworkspaces

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnConnectionAliasMixinProps",
		reflect.TypeOf((*CfnConnectionAliasMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnConnectionAliasPropsMixin",
		reflect.TypeOf((*CfnConnectionAliasPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectionAliasPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnConnectionAliasPropsMixin.ConnectionAliasAssociationProperty",
		reflect.TypeOf((*CfnConnectionAliasPropsMixin_ConnectionAliasAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspaceMixinProps",
		reflect.TypeOf((*CfnWorkspaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspacePropsMixin",
		reflect.TypeOf((*CfnWorkspacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkspacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspacePropsMixin.WorkspacePropertiesProperty",
		reflect.TypeOf((*CfnWorkspacePropsMixin_WorkspacePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspacesPoolMixinProps",
		reflect.TypeOf((*CfnWorkspacesPoolMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspacesPoolPropsMixin",
		reflect.TypeOf((*CfnWorkspacesPoolPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkspacesPoolPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspacesPoolPropsMixin.ApplicationSettingsProperty",
		reflect.TypeOf((*CfnWorkspacesPoolPropsMixin_ApplicationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspacesPoolPropsMixin.CapacityProperty",
		reflect.TypeOf((*CfnWorkspacesPoolPropsMixin_CapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspacesPoolPropsMixin.TimeoutSettingsProperty",
		reflect.TypeOf((*CfnWorkspacesPoolPropsMixin_TimeoutSettingsProperty)(nil)).Elem(),
	)
}
