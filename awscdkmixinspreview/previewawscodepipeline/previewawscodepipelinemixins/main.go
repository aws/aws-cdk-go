package previewawscodepipelinemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnCustomActionTypeMixinProps",
		reflect.TypeOf((*CfnCustomActionTypeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnCustomActionTypePropsMixin",
		reflect.TypeOf((*CfnCustomActionTypePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomActionTypePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnCustomActionTypePropsMixin.ArtifactDetailsProperty",
		reflect.TypeOf((*CfnCustomActionTypePropsMixin_ArtifactDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnCustomActionTypePropsMixin.ConfigurationPropertiesProperty",
		reflect.TypeOf((*CfnCustomActionTypePropsMixin_ConfigurationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnCustomActionTypePropsMixin.SettingsProperty",
		reflect.TypeOf((*CfnCustomActionTypePropsMixin_SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelineMixinProps",
		reflect.TypeOf((*CfnPipelineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin",
		reflect.TypeOf((*CfnPipelinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipelinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.ActionDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ActionDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.ActionTypeIdProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ActionTypeIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.ArtifactStoreMapProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ArtifactStoreMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.ArtifactStoreProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ArtifactStoreProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.BeforeEntryConditionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_BeforeEntryConditionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.BlockerDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_BlockerDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.EncryptionKeyProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_EncryptionKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.EnvironmentVariableProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_EnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.FailureConditionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_FailureConditionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.GitBranchFilterCriteriaProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitBranchFilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.GitConfigurationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.GitFilePathFilterCriteriaProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitFilePathFilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.GitPullRequestFilterProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitPullRequestFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.GitPushFilterProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitPushFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.GitTagFilterCriteriaProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitTagFilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.InputArtifactProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_InputArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.OutputArtifactProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_OutputArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.PipelineTriggerDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_PipelineTriggerDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.RetryConfigurationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_RetryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.RuleDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_RuleDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.RuleTypeIdProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_RuleTypeIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.StageDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_StageDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.StageTransitionProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_StageTransitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.SuccessConditionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_SuccessConditionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin.VariableDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_VariableDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnWebhookMixinProps",
		reflect.TypeOf((*CfnWebhookMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnWebhookPropsMixin",
		reflect.TypeOf((*CfnWebhookPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWebhookPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnWebhookPropsMixin.WebhookAuthConfigurationProperty",
		reflect.TypeOf((*CfnWebhookPropsMixin_WebhookAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnWebhookPropsMixin.WebhookFilterRuleProperty",
		reflect.TypeOf((*CfnWebhookPropsMixin_WebhookFilterRuleProperty)(nil)).Elem(),
	)
}
