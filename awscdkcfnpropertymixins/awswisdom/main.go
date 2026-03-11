package awswisdom

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentMixinProps",
		reflect.TypeOf((*CfnAIAgentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin",
		reflect.TypeOf((*CfnAIAgentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAIAgentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.AIAgentConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_AIAgentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.AnswerRecommendationAIAgentConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_AnswerRecommendationAIAgentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.AssociationConfigurationDataProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_AssociationConfigurationDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.AssociationConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_AssociationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.CaseSummarizationAIAgentConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_CaseSummarizationAIAgentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.EmailGenerativeAnswerAIAgentConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_EmailGenerativeAnswerAIAgentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.EmailOverviewAIAgentConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_EmailOverviewAIAgentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.EmailResponseAIAgentConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_EmailResponseAIAgentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.KnowledgeBaseAssociationConfigurationDataProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_KnowledgeBaseAssociationConfigurationDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.ManualSearchAIAgentConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_ManualSearchAIAgentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.NoteTakingAIAgentConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_NoteTakingAIAgentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.OrConditionProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_OrConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.OrchestrationAIAgentConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_OrchestrationAIAgentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.SelfServiceAIAgentConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_SelfServiceAIAgentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.TagConditionProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_TagConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.TagFilterProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.ToolConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_ToolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.ToolInstructionProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_ToolInstructionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.ToolOutputConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_ToolOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.ToolOutputFilterProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_ToolOutputFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.ToolOverrideConstantInputValueProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_ToolOverrideConstantInputValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.ToolOverrideInputValueConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_ToolOverrideInputValueConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.ToolOverrideInputValueProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_ToolOverrideInputValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentPropsMixin.UserInteractionConfigurationProperty",
		reflect.TypeOf((*CfnAIAgentPropsMixin_UserInteractionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentVersionMixinProps",
		reflect.TypeOf((*CfnAIAgentVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIAgentVersionPropsMixin",
		reflect.TypeOf((*CfnAIAgentVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAIAgentVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailMixinProps",
		reflect.TypeOf((*CfnAIGuardrailMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAIGuardrailPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.AIGuardrailContentPolicyConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_AIGuardrailContentPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.AIGuardrailContextualGroundingPolicyConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_AIGuardrailContextualGroundingPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.AIGuardrailSensitiveInformationPolicyConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_AIGuardrailSensitiveInformationPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.AIGuardrailTopicPolicyConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_AIGuardrailTopicPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.AIGuardrailWordPolicyConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_AIGuardrailWordPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.GuardrailContentFilterConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_GuardrailContentFilterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.GuardrailContextualGroundingFilterConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_GuardrailContextualGroundingFilterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.GuardrailManagedWordsConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_GuardrailManagedWordsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.GuardrailPiiEntityConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_GuardrailPiiEntityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.GuardrailRegexConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_GuardrailRegexConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.GuardrailTopicConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_GuardrailTopicConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailPropsMixin.GuardrailWordConfigProperty",
		reflect.TypeOf((*CfnAIGuardrailPropsMixin_GuardrailWordConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailVersionMixinProps",
		reflect.TypeOf((*CfnAIGuardrailVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIGuardrailVersionPropsMixin",
		reflect.TypeOf((*CfnAIGuardrailVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAIGuardrailVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIPromptMixinProps",
		reflect.TypeOf((*CfnAIPromptMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIPromptPropsMixin",
		reflect.TypeOf((*CfnAIPromptPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAIPromptPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIPromptPropsMixin.AIPromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnAIPromptPropsMixin_AIPromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIPromptPropsMixin.TextFullAIPromptEditTemplateConfigurationProperty",
		reflect.TypeOf((*CfnAIPromptPropsMixin_TextFullAIPromptEditTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIPromptVersionMixinProps",
		reflect.TypeOf((*CfnAIPromptVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIPromptVersionPropsMixin",
		reflect.TypeOf((*CfnAIPromptVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAIPromptVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantAssociationMixinProps",
		reflect.TypeOf((*CfnAssistantAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantAssociationPropsMixin",
		reflect.TypeOf((*CfnAssistantAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssistantAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantAssociationPropsMixin.AssociationDataProperty",
		reflect.TypeOf((*CfnAssistantAssociationPropsMixin_AssociationDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantAssociationPropsMixin.ExternalBedrockKnowledgeBaseConfigProperty",
		reflect.TypeOf((*CfnAssistantAssociationPropsMixin_ExternalBedrockKnowledgeBaseConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantMixinProps",
		reflect.TypeOf((*CfnAssistantMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantPropsMixin",
		reflect.TypeOf((*CfnAssistantPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssistantPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAssistantPropsMixin.ServerSideEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnAssistantPropsMixin_ServerSideEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBaseMixinProps",
		reflect.TypeOf((*CfnKnowledgeBaseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKnowledgeBasePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.AppIntegrationsConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_AppIntegrationsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.BedrockFoundationModelConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_BedrockFoundationModelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.ChunkingConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_ChunkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.CrawlerLimitsProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_CrawlerLimitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.FixedSizeChunkingConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_FixedSizeChunkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.HierarchicalChunkingConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_HierarchicalChunkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.HierarchicalChunkingLevelConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_HierarchicalChunkingLevelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.ManagedSourceConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_ManagedSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.ParsingConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_ParsingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.ParsingPromptProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_ParsingPromptProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.RenderingConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_RenderingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.SeedUrlProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_SeedUrlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.SemanticChunkingConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_SemanticChunkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.ServerSideEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_ServerSideEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.SourceConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_SourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.UrlConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_UrlConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.VectorIngestionConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_VectorIngestionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnKnowledgeBasePropsMixin.WebCrawlerConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBasePropsMixin_WebCrawlerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplateMixinProps",
		reflect.TypeOf((*CfnMessageTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMessageTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.AgentAttributesProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_AgentAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.ContentProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_ContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.CustomerProfileAttributesProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_CustomerProfileAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.EmailMessageTemplateContentBodyProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_EmailMessageTemplateContentBodyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.EmailMessageTemplateContentProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_EmailMessageTemplateContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.EmailMessageTemplateHeaderProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_EmailMessageTemplateHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.GroupingConfigurationProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_GroupingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.MessageTemplateAttachmentProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_MessageTemplateAttachmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.MessageTemplateAttributesProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_MessageTemplateAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.MessageTemplateBodyContentProviderProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_MessageTemplateBodyContentProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.SmsMessageTemplateContentBodyProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_SmsMessageTemplateContentBodyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.SmsMessageTemplateContentProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_SmsMessageTemplateContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.SystemAttributesProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_SystemAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplatePropsMixin.SystemEndpointAttributesProperty",
		reflect.TypeOf((*CfnMessageTemplatePropsMixin_SystemEndpointAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplateVersionMixinProps",
		reflect.TypeOf((*CfnMessageTemplateVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnMessageTemplateVersionPropsMixin",
		reflect.TypeOf((*CfnMessageTemplateVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMessageTemplateVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnQuickResponseMixinProps",
		reflect.TypeOf((*CfnQuickResponseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnQuickResponsePropsMixin",
		reflect.TypeOf((*CfnQuickResponsePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQuickResponsePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnQuickResponsePropsMixin.GroupingConfigurationProperty",
		reflect.TypeOf((*CfnQuickResponsePropsMixin_GroupingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnQuickResponsePropsMixin.QuickResponseContentProviderProperty",
		reflect.TypeOf((*CfnQuickResponsePropsMixin_QuickResponseContentProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnQuickResponsePropsMixin.QuickResponseContentsProperty",
		reflect.TypeOf((*CfnQuickResponsePropsMixin_QuickResponseContentsProperty)(nil)).Elem(),
	)
}
