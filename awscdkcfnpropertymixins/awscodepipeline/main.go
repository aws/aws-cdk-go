package awscodepipeline

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnCustomActionTypeMixinProps",
		reflect.TypeOf((*CfnCustomActionTypeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnCustomActionTypePropsMixin",
		reflect.TypeOf((*CfnCustomActionTypePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomActionTypePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnCustomActionTypePropsMixin.ArtifactDetailsProperty",
		reflect.TypeOf((*CfnCustomActionTypePropsMixin_ArtifactDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnCustomActionTypePropsMixin.ConfigurationPropertiesProperty",
		reflect.TypeOf((*CfnCustomActionTypePropsMixin_ConfigurationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnCustomActionTypePropsMixin.SettingsProperty",
		reflect.TypeOf((*CfnCustomActionTypePropsMixin_SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelineMixinProps",
		reflect.TypeOf((*CfnPipelineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin",
		reflect.TypeOf((*CfnPipelinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipelinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.ActionDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ActionDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.ActionTypeIdProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ActionTypeIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.ArtifactStoreMapProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ArtifactStoreMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.ArtifactStoreProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ArtifactStoreProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.BeforeEntryConditionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_BeforeEntryConditionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.BlockerDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_BlockerDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.EncryptionKeyProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_EncryptionKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.EnvironmentVariableProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_EnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.FailureConditionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_FailureConditionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.GitBranchFilterCriteriaProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitBranchFilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.GitConfigurationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.GitFilePathFilterCriteriaProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitFilePathFilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.GitPullRequestFilterProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitPullRequestFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.GitPushFilterProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitPushFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.GitTagFilterCriteriaProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_GitTagFilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.InputArtifactProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_InputArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.OutputArtifactProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_OutputArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.PipelineTriggerDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_PipelineTriggerDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.RetryConfigurationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_RetryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.RuleDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_RuleDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.RuleTypeIdProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_RuleTypeIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.StageDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_StageDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.StageTransitionProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_StageTransitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.SuccessConditionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_SuccessConditionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnPipelinePropsMixin.VariableDeclarationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_VariableDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnWebhookMixinProps",
		reflect.TypeOf((*CfnWebhookMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnWebhookPropsMixin",
		reflect.TypeOf((*CfnWebhookPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWebhookPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnWebhookPropsMixin.WebhookAuthConfigurationProperty",
		reflect.TypeOf((*CfnWebhookPropsMixin_WebhookAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codepipeline.CfnWebhookPropsMixin.WebhookFilterRuleProperty",
		reflect.TypeOf((*CfnWebhookPropsMixin_WebhookFilterRuleProperty)(nil)).Elem(),
	)
}
