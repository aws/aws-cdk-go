package awsresiliencehubv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnPolicyMixinProps",
		reflect.TypeOf((*CfnPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnPolicyPropsMixin",
		reflect.TypeOf((*CfnPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnPolicyPropsMixin.AvailabilitySloProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_AvailabilitySloProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnPolicyPropsMixin.DataRecoveryTargetsProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_DataRecoveryTargetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnPolicyPropsMixin.MultiAzTargetsProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_MultiAzTargetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnPolicyPropsMixin.MultiRegionTargetsProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_MultiRegionTargetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServiceFunctionMixinProps",
		reflect.TypeOf((*CfnServiceFunctionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServiceFunctionPropsMixin",
		reflect.TypeOf((*CfnServiceFunctionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServiceFunctionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServiceMixinProps",
		reflect.TypeOf((*CfnServiceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin",
		reflect.TypeOf((*CfnServicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.AssertionDefinitionProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AssertionDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.AssociatedSystemProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AssociatedSystemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.CrossAccountRoleConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CrossAccountRoleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.DisasterRecoverySourceProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DisasterRecoverySourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.EffectivePolicyValuesProperty",
		reflect.TypeOf((*CfnServicePropsMixin_EffectivePolicyValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.EksSourceProperty",
		reflect.TypeOf((*CfnServicePropsMixin_EksSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.InputSourceDefinitionProperty",
		reflect.TypeOf((*CfnServicePropsMixin_InputSourceDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.PermissionModelProperty",
		reflect.TypeOf((*CfnServicePropsMixin_PermissionModelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.ReportOutputConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ReportOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.ResourceConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ResourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.ResourceTagProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ResourceTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.S3ReportOutputConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_S3ReportOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.ServiceReportConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceReportConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.SloSourceProperty",
		reflect.TypeOf((*CfnServicePropsMixin_SloSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin.TargetSourceProperty",
		reflect.TypeOf((*CfnServicePropsMixin_TargetSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnSystemMixinProps",
		reflect.TypeOf((*CfnSystemMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnSystemPropsMixin",
		reflect.TypeOf((*CfnSystemPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSystemPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnUserJourneyMixinProps",
		reflect.TypeOf((*CfnUserJourneyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnUserJourneyPropsMixin",
		reflect.TypeOf((*CfnUserJourneyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserJourneyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
