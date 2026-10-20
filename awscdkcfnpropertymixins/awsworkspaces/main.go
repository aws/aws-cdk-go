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
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryMixinProps",
		reflect.TypeOf((*CfnDirectoryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin",
		reflect.TypeOf((*CfnDirectoryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDirectoryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.AccessEndpointConfigProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_AccessEndpointConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.AccessEndpointProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_AccessEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.ActiveDirectoryConfigProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_ActiveDirectoryConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.CertificateBasedAuthPropertiesProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_CertificateBasedAuthPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.DefaultWorkspaceCreationPropertiesProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_DefaultWorkspaceCreationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.GlobalAcceleratorForDirectoryProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_GlobalAcceleratorForDirectoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.IDCConfigProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_IDCConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.MicrosoftEntraConfigProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_MicrosoftEntraConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.SamlPropertiesProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_SamlPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.SelfservicePermissionsProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_SelfservicePermissionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.StorageConnectorProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_StorageConnectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.StreamingPropertiesProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_StreamingPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.UserSettingProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_UserSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin.WorkspaceAccessPropertiesProperty",
		reflect.TypeOf((*CfnDirectoryPropsMixin_WorkspaceAccessPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspaceIpGroupMixinProps",
		reflect.TypeOf((*CfnWorkspaceIpGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspaceIpGroupPropsMixin",
		reflect.TypeOf((*CfnWorkspaceIpGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkspaceIpGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnWorkspaceIpGroupPropsMixin.IpRuleItemProperty",
		reflect.TypeOf((*CfnWorkspaceIpGroupPropsMixin_IpRuleItemProperty)(nil)).Elem(),
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
