package awsiotsitewise

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAccessPolicyMixinProps",
		reflect.TypeOf((*CfnAccessPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAccessPolicyPropsMixin",
		reflect.TypeOf((*CfnAccessPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAccessPolicyPropsMixin.AccessPolicyIdentityProperty",
		reflect.TypeOf((*CfnAccessPolicyPropsMixin_AccessPolicyIdentityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAccessPolicyPropsMixin.AccessPolicyResourceProperty",
		reflect.TypeOf((*CfnAccessPolicyPropsMixin_AccessPolicyResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAccessPolicyPropsMixin.IamRoleProperty",
		reflect.TypeOf((*CfnAccessPolicyPropsMixin_IamRoleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAccessPolicyPropsMixin.IamUserProperty",
		reflect.TypeOf((*CfnAccessPolicyPropsMixin_IamUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAccessPolicyPropsMixin.PortalProperty",
		reflect.TypeOf((*CfnAccessPolicyPropsMixin_PortalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAccessPolicyPropsMixin.ProjectProperty",
		reflect.TypeOf((*CfnAccessPolicyPropsMixin_ProjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAccessPolicyPropsMixin.UserProperty",
		reflect.TypeOf((*CfnAccessPolicyPropsMixin_UserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetMixinProps",
		reflect.TypeOf((*CfnAssetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelMixinProps",
		reflect.TypeOf((*CfnAssetModelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin",
		reflect.TypeOf((*CfnAssetModelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssetModelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.AssetModelCompositeModelProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_AssetModelCompositeModelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.AssetModelHierarchyProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_AssetModelHierarchyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.AssetModelPropertyProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_AssetModelPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.AttributeProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_AttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.EnforcedAssetModelInterfacePropertyMappingProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_EnforcedAssetModelInterfacePropertyMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.EnforcedAssetModelInterfaceRelationshipProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_EnforcedAssetModelInterfaceRelationshipProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.ExpressionVariableProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_ExpressionVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.MetricProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.MetricWindowProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_MetricWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.PropertyPathDefinitionProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_PropertyPathDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.PropertyTypeProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_PropertyTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.TransformProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_TransformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.TumblingWindowProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_TumblingWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetModelPropsMixin.VariableValueProperty",
		reflect.TypeOf((*CfnAssetModelPropsMixin_VariableValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetPropsMixin",
		reflect.TypeOf((*CfnAssetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetPropsMixin.AssetHierarchyProperty",
		reflect.TypeOf((*CfnAssetPropsMixin_AssetHierarchyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnAssetPropsMixin.AssetPropertyProperty",
		reflect.TypeOf((*CfnAssetPropsMixin_AssetPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnComputationModelMixinProps",
		reflect.TypeOf((*CfnComputationModelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnComputationModelPropsMixin",
		reflect.TypeOf((*CfnComputationModelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComputationModelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnComputationModelPropsMixin.AnomalyDetectionComputationModelConfigurationProperty",
		reflect.TypeOf((*CfnComputationModelPropsMixin_AnomalyDetectionComputationModelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnComputationModelPropsMixin.AssetModelPropertyBindingValueProperty",
		reflect.TypeOf((*CfnComputationModelPropsMixin_AssetModelPropertyBindingValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnComputationModelPropsMixin.AssetPropertyBindingValueProperty",
		reflect.TypeOf((*CfnComputationModelPropsMixin_AssetPropertyBindingValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnComputationModelPropsMixin.ComputationModelConfigurationProperty",
		reflect.TypeOf((*CfnComputationModelPropsMixin_ComputationModelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnComputationModelPropsMixin.ComputationModelDataBindingValueProperty",
		reflect.TypeOf((*CfnComputationModelPropsMixin_ComputationModelDataBindingValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnDashboardMixinProps",
		reflect.TypeOf((*CfnDashboardMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnDashboardPropsMixin",
		reflect.TypeOf((*CfnDashboardPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDashboardPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnDatasetMixinProps",
		reflect.TypeOf((*CfnDatasetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnDatasetPropsMixin",
		reflect.TypeOf((*CfnDatasetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatasetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnDatasetPropsMixin.DatasetSourceProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DatasetSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnDatasetPropsMixin.KendraSourceDetailProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_KendraSourceDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnDatasetPropsMixin.SourceDetailProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_SourceDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnGatewayMixinProps",
		reflect.TypeOf((*CfnGatewayMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnGatewayPropsMixin",
		reflect.TypeOf((*CfnGatewayPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnGatewayPropsMixin.GatewayCapabilitySummaryProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GatewayCapabilitySummaryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnGatewayPropsMixin.GatewayPlatformProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GatewayPlatformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnGatewayPropsMixin.GreengrassProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GreengrassProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnGatewayPropsMixin.GreengrassV2Property",
		reflect.TypeOf((*CfnGatewayPropsMixin_GreengrassV2Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnGatewayPropsMixin.SiemensIEProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_SiemensIEProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnPortalMixinProps",
		reflect.TypeOf((*CfnPortalMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnPortalPropsMixin",
		reflect.TypeOf((*CfnPortalPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPortalPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnPortalPropsMixin.AlarmsProperty",
		reflect.TypeOf((*CfnPortalPropsMixin_AlarmsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnPortalPropsMixin.PortalTypeEntryProperty",
		reflect.TypeOf((*CfnPortalPropsMixin_PortalTypeEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
