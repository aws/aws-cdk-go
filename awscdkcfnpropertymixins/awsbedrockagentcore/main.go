package awsbedrockagentcore

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnApiKeyCredentialProviderMixinProps",
		reflect.TypeOf((*CfnApiKeyCredentialProviderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnApiKeyCredentialProviderPropsMixin",
		reflect.TypeOf((*CfnApiKeyCredentialProviderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApiKeyCredentialProviderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnApiKeyCredentialProviderPropsMixin.ApiKeySecretArnProperty",
		reflect.TypeOf((*CfnApiKeyCredentialProviderPropsMixin_ApiKeySecretArnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnBrowserCustomMixinProps",
		reflect.TypeOf((*CfnBrowserCustomMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnBrowserCustomPropsMixin",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrowserCustomPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnBrowserCustomPropsMixin.BrowserNetworkConfigurationProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_BrowserNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnBrowserCustomPropsMixin.BrowserSigningProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_BrowserSigningProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnBrowserCustomPropsMixin.RecordingConfigProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_RecordingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnBrowserCustomPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnBrowserCustomPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnBrowserProfileMixinProps",
		reflect.TypeOf((*CfnBrowserProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnBrowserProfilePropsMixin",
		reflect.TypeOf((*CfnBrowserProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrowserProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnCodeInterpreterCustomMixinProps",
		reflect.TypeOf((*CfnCodeInterpreterCustomMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnCodeInterpreterCustomPropsMixin",
		reflect.TypeOf((*CfnCodeInterpreterCustomPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeInterpreterCustomPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnCodeInterpreterCustomPropsMixin.CodeInterpreterNetworkConfigurationProperty",
		reflect.TypeOf((*CfnCodeInterpreterCustomPropsMixin_CodeInterpreterNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnCodeInterpreterCustomPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnCodeInterpreterCustomPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorMixinProps",
		reflect.TypeOf((*CfnEvaluatorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin",
		reflect.TypeOf((*CfnEvaluatorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEvaluatorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin.BedrockEvaluatorModelConfigProperty",
		reflect.TypeOf((*CfnEvaluatorPropsMixin_BedrockEvaluatorModelConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin.CategoricalScaleDefinitionProperty",
		reflect.TypeOf((*CfnEvaluatorPropsMixin_CategoricalScaleDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin.CodeBasedEvaluatorConfigProperty",
		reflect.TypeOf((*CfnEvaluatorPropsMixin_CodeBasedEvaluatorConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin.EvaluatorConfigProperty",
		reflect.TypeOf((*CfnEvaluatorPropsMixin_EvaluatorConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin.EvaluatorModelConfigProperty",
		reflect.TypeOf((*CfnEvaluatorPropsMixin_EvaluatorModelConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin.InferenceConfigurationProperty",
		reflect.TypeOf((*CfnEvaluatorPropsMixin_InferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin.LambdaEvaluatorConfigProperty",
		reflect.TypeOf((*CfnEvaluatorPropsMixin_LambdaEvaluatorConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin.LlmAsAJudgeEvaluatorConfigProperty",
		reflect.TypeOf((*CfnEvaluatorPropsMixin_LlmAsAJudgeEvaluatorConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin.NumericalScaleDefinitionProperty",
		reflect.TypeOf((*CfnEvaluatorPropsMixin_NumericalScaleDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin.RatingScaleProperty",
		reflect.TypeOf((*CfnEvaluatorPropsMixin_RatingScaleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayMixinProps",
		reflect.TypeOf((*CfnGatewayMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.AuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_AuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.AuthorizingClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_AuthorizingClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.ClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_ClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.CustomClaimValidationTypeProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_CustomClaimValidationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.CustomJWTAuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_CustomJWTAuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.GatewayInterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GatewayInterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.GatewayPolicyEngineConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GatewayPolicyEngineConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.GatewayProtocolConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GatewayProtocolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.InterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_InterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.InterceptorInputConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_InterceptorInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.LambdaInterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_LambdaInterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.MCPGatewayConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_MCPGatewayConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayPropsMixin.WorkloadIdentityDetailsProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_WorkloadIdentityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetMixinProps",
		reflect.TypeOf((*CfnGatewayTargetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayTargetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.ApiGatewayTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiGatewayTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.ApiGatewayToolConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiGatewayToolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.ApiGatewayToolFilterProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiGatewayToolFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.ApiGatewayToolOverrideProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiGatewayToolOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.ApiKeyCredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiKeyCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.ApiSchemaConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiSchemaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.CredentialProviderConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_CredentialProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.CredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_CredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.IamCredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_IamCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.McpLambdaTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_McpLambdaTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.McpServerTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_McpServerTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.McpTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_McpTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.MetadataConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_MetadataConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.OAuthCredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_OAuthCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.SchemaDefinitionProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_SchemaDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.TargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_TargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.ToolDefinitionProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ToolDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnGatewayTargetPropsMixin.ToolSchemaProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ToolSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryMixinProps",
		reflect.TypeOf((*CfnMemoryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin",
		reflect.TypeOf((*CfnMemoryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMemoryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.ContentConfigurationProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_ContentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.CustomConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_CustomConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.CustomMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_CustomMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.EpisodicMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.EpisodicOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.EpisodicOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.EpisodicOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.EpisodicOverrideReflectionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideReflectionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.EpisodicReflectionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicReflectionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.InvocationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_InvocationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.KinesisResourceProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_KinesisResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.MemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_MemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.MessageBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_MessageBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.SelfManagedConfigurationProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SelfManagedConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.SemanticMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.SemanticOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.SemanticOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.SemanticOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.StreamDeliveryResourceProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_StreamDeliveryResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.StreamDeliveryResourcesProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_StreamDeliveryResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.SummaryMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SummaryMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.SummaryOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SummaryOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.SummaryOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SummaryOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.TimeBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_TimeBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.TokenBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_TokenBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.TriggerConditionInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_TriggerConditionInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.UserPreferenceMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.UserPreferenceOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.UserPreferenceOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin.UserPreferenceOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderMixinProps",
		reflect.TypeOf((*CfnOAuth2CredentialProviderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOAuth2CredentialProviderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.AtlassianOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_AtlassianOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.ClientSecretArnProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_ClientSecretArnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.CustomOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_CustomOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.GithubOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_GithubOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.GoogleOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_GoogleOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.IncludedOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_IncludedOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.LinkedinOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_LinkedinOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.MicrosoftOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_MicrosoftOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.Oauth2AuthorizationServerMetadataProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_Oauth2AuthorizationServerMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.Oauth2DiscoveryProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_Oauth2DiscoveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.Oauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_Oauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.Oauth2ProviderConfigOutputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_Oauth2ProviderConfigOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.SalesforceOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_SalesforceOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin.SlackOauth2ProviderConfigInputProperty",
		reflect.TypeOf((*CfnOAuth2CredentialProviderPropsMixin_SlackOauth2ProviderConfigInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigMixinProps",
		reflect.TypeOf((*CfnOnlineEvaluationConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOnlineEvaluationConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin.CloudWatchLogsInputConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin_CloudWatchLogsInputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin.CloudWatchOutputConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin_CloudWatchOutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin.DataSourceConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin_DataSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin.EvaluatorReferenceProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin_EvaluatorReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin.FilterValueProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin_FilterValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin.OutputConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin_OutputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin.SamplingConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin_SamplingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOnlineEvaluationConfigPropsMixin.SessionConfigProperty",
		reflect.TypeOf((*CfnOnlineEvaluationConfigPropsMixin_SessionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPolicyEngineMixinProps",
		reflect.TypeOf((*CfnPolicyEngineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPolicyEnginePropsMixin",
		reflect.TypeOf((*CfnPolicyEnginePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyEnginePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPolicyMixinProps",
		reflect.TypeOf((*CfnPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPolicyPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPolicyPropsMixin.CedarPolicyProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_CedarPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPolicyPropsMixin.PolicyDefinitionProperty",
		reflect.TypeOf((*CfnPolicyPropsMixin_PolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimeEndpointMixinProps",
		reflect.TypeOf((*CfnRuntimeEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimeEndpointPropsMixin",
		reflect.TypeOf((*CfnRuntimeEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuntimeEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimeMixinProps",
		reflect.TypeOf((*CfnRuntimeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin",
		reflect.TypeOf((*CfnRuntimePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuntimePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.AgentRuntimeArtifactProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_AgentRuntimeArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.AuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_AuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.AuthorizingClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_AuthorizingClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.ClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_ClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.CodeConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.CodeProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.ContainerConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_ContainerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.CustomClaimValidationTypeProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CustomClaimValidationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.CustomJWTAuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CustomJWTAuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.FilesystemConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_FilesystemConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.LifecycleConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_LifecycleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.RequestHeaderConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_RequestHeaderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.SessionStorageConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_SessionStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnRuntimePropsMixin.WorkloadIdentityDetailsProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_WorkloadIdentityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnWorkloadIdentityMixinProps",
		reflect.TypeOf((*CfnWorkloadIdentityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnWorkloadIdentityPropsMixin",
		reflect.TypeOf((*CfnWorkloadIdentityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkloadIdentityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
