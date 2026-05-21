package awsentityresolution

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdMappingWorkflowMixinProps",
		reflect.TypeOf((*CfnIdMappingWorkflowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdMappingWorkflowPropsMixin",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdMappingWorkflowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdMappingWorkflowPropsMixin.IdMappingIncrementalRunConfigProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IdMappingIncrementalRunConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdMappingWorkflowPropsMixin.IdMappingRuleBasedPropertiesProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IdMappingRuleBasedPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdMappingWorkflowPropsMixin.IdMappingTechniquesProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IdMappingTechniquesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdMappingWorkflowPropsMixin.IdMappingWorkflowInputSourceProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IdMappingWorkflowInputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdMappingWorkflowPropsMixin.IdMappingWorkflowOutputSourceProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IdMappingWorkflowOutputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdMappingWorkflowPropsMixin.IntermediateSourceConfigurationProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IntermediateSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdMappingWorkflowPropsMixin.ProviderPropertiesProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_ProviderPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdMappingWorkflowPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdNamespaceMixinProps",
		reflect.TypeOf((*CfnIdNamespaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdNamespacePropsMixin",
		reflect.TypeOf((*CfnIdNamespacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdNamespacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdNamespacePropsMixin.IdNamespaceIdMappingWorkflowPropertiesProperty",
		reflect.TypeOf((*CfnIdNamespacePropsMixin_IdNamespaceIdMappingWorkflowPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdNamespacePropsMixin.IdNamespaceInputSourceProperty",
		reflect.TypeOf((*CfnIdNamespacePropsMixin_IdNamespaceInputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdNamespacePropsMixin.NamespaceProviderPropertiesProperty",
		reflect.TypeOf((*CfnIdNamespacePropsMixin_NamespaceProviderPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdNamespacePropsMixin.NamespaceRuleBasedPropertiesProperty",
		reflect.TypeOf((*CfnIdNamespacePropsMixin_NamespaceRuleBasedPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnIdNamespacePropsMixin.RuleProperty",
		reflect.TypeOf((*CfnIdNamespacePropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowMixinProps",
		reflect.TypeOf((*CfnMatchingWorkflowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMatchingWorkflowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.CustomerProfilesIntegrationConfigProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_CustomerProfilesIntegrationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.IncrementalRunConfigProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_IncrementalRunConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.InputSourceProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_InputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.IntermediateSourceConfigurationProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_IntermediateSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.MatchingConfigProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_MatchingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.OutputAttributeProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_OutputAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.OutputSourceProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_OutputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.ProviderPropertiesProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_ProviderPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.ResolutionTechniquesProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_ResolutionTechniquesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.RuleBasedPropertiesProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_RuleBasedPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.RuleConditionPropertiesProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_RuleConditionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.RuleConditionProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_RuleConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnMatchingWorkflowPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnPolicyStatementMixinProps",
		reflect.TypeOf((*CfnPolicyStatementMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnPolicyStatementPropsMixin",
		reflect.TypeOf((*CfnPolicyStatementPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyStatementPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnSchemaMappingMixinProps",
		reflect.TypeOf((*CfnSchemaMappingMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnSchemaMappingPropsMixin",
		reflect.TypeOf((*CfnSchemaMappingPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchemaMappingPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_entityresolution.CfnSchemaMappingPropsMixin.SchemaInputAttributeProperty",
		reflect.TypeOf((*CfnSchemaMappingPropsMixin_SchemaInputAttributeProperty)(nil)).Elem(),
	)
}
