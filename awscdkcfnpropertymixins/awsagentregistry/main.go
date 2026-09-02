package awsagentregistry

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryMixinProps",
		reflect.TypeOf((*CfnRegistryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryPropsMixin",
		reflect.TypeOf((*CfnRegistryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRegistryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryPropsMixin.ApprovalConfigurationProperty",
		reflect.TypeOf((*CfnRegistryPropsMixin_ApprovalConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryPropsMixin.AuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnRegistryPropsMixin_AuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryPropsMixin.AuthorizingClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnRegistryPropsMixin_AuthorizingClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryPropsMixin.ClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnRegistryPropsMixin_ClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryPropsMixin.CustomClaimValidationTypeProperty",
		reflect.TypeOf((*CfnRegistryPropsMixin_CustomClaimValidationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryPropsMixin.CustomJWTAuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnRegistryPropsMixin_CustomJWTAuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryPropsMixin.DiscoveryConfigurationProperty",
		reflect.TypeOf((*CfnRegistryPropsMixin_DiscoveryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordMixinProps",
		reflect.TypeOf((*CfnRegistryRecordMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRegistryRecordPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.A2aAgentCardDescriptorProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_A2aAgentCardDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.AgentSkillsAdditionalDataProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_AgentSkillsAdditionalDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.AgentSkillsDefinitionDescriptorProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_AgentSkillsDefinitionDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.AgentSkillsMdDescriptorProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_AgentSkillsMdDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.CustomDescriptorProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_CustomDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.DescriptorSourceFromUrlProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_DescriptorSourceFromUrlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.DescriptorSourceProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_DescriptorSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.DescriptorsProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_DescriptorsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.McpServerAdditionalDataProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_McpServerAdditionalDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.McpServerDescriptorProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_McpServerDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.McpToolsDescriptorProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_McpToolsDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.RegistryRecordCredentialProviderConfigurationProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_RegistryRecordCredentialProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.RegistryRecordCredentialProviderUnionProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_RegistryRecordCredentialProviderUnionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.RegistryRecordIamCredentialProviderProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_RegistryRecordIamCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.RegistryRecordOAuthCredentialProviderProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_RegistryRecordOAuthCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.SkillMdSourceFromUrlProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_SkillMdSourceFromUrlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_agentregistry.CfnRegistryRecordPropsMixin.SkillMdSourceProperty",
		reflect.TypeOf((*CfnRegistryRecordPropsMixin_SkillMdSourceProperty)(nil)).Elem(),
	)
}
