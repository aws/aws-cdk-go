package previewawsbedrockagentcoremixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomLogsMixin",
		reflect.TypeOf((*CfnBrowserCustomLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrowserCustomLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomMixinProps",
		reflect.TypeOf((*CfnBrowserCustomMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrowserCustomPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin.BrowserNetworkConfigurationProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_BrowserNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin.BrowserSigningProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_BrowserSigningProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin.RecordingConfigProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_RecordingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogs",
		reflect.TypeOf((*CfnBrowserCustomUsageLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnBrowserCustomUsageLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogs",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCodeInterpreterCustomApplicationLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomLogsMixin",
		reflect.TypeOf((*CfnCodeInterpreterCustomLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeInterpreterCustomLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomMixinProps",
		reflect.TypeOf((*CfnCodeInterpreterCustomMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomPropsMixin",
		reflect.TypeOf((*CfnCodeInterpreterCustomPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeInterpreterCustomPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomPropsMixin.CodeInterpreterNetworkConfigurationProperty",
		reflect.TypeOf((*CfnCodeInterpreterCustomPropsMixin_CodeInterpreterNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnCodeInterpreterCustomPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogs",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCodeInterpreterCustomUsageLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogs",
		reflect.TypeOf((*CfnGatewayApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnGatewayApplicationLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayLogsMixin",
		reflect.TypeOf((*CfnGatewayLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayMixinProps",
		reflect.TypeOf((*CfnGatewayMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin",
		reflect.TypeOf((*CfnGatewayPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.AuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_AuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.AuthorizingClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_AuthorizingClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.ClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_ClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.CustomClaimValidationTypeProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_CustomClaimValidationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.CustomJWTAuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_CustomJWTAuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.GatewayInterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GatewayInterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.GatewayProtocolConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GatewayProtocolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.InterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_InterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.InterceptorInputConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_InterceptorInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.LambdaInterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_LambdaInterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.MCPGatewayConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_MCPGatewayConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.WorkloadIdentityDetailsProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_WorkloadIdentityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetMixinProps",
		reflect.TypeOf((*CfnGatewayTargetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayTargetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ApiKeyCredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiKeyCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ApiSchemaConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiSchemaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.CredentialProviderConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_CredentialProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.CredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_CredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.McpLambdaTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_McpLambdaTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.McpServerTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_McpServerTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.McpTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_McpTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.MetadataConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_MetadataConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.OAuthCredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_OAuthCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.SchemaDefinitionProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_SchemaDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.TargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_TargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ToolDefinitionProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ToolDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ToolSchemaProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ToolSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTraces",
		reflect.TypeOf((*CfnGatewayTraces)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toXRay", GoMethod: "ToXRay"},
		},
		func() interface{} {
			return &jsiiProxy_CfnGatewayTraces{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogs",
		reflect.TypeOf((*CfnMemoryApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMemoryApplicationLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryLogsMixin",
		reflect.TypeOf((*CfnMemoryLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMemoryLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryMixinProps",
		reflect.TypeOf((*CfnMemoryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin",
		reflect.TypeOf((*CfnMemoryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMemoryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.CustomConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_CustomConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.CustomMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_CustomMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicOverrideReflectionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideReflectionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicReflectionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicReflectionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.InvocationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_InvocationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.MemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_MemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.MessageBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_MessageBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SelfManagedConfigurationProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SelfManagedConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SemanticMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SemanticOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SemanticOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SemanticOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SummaryMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SummaryMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SummaryOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SummaryOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SummaryOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SummaryOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.TimeBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_TimeBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.TokenBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_TokenBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.TriggerConditionInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_TriggerConditionInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.UserPreferenceMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.UserPreferenceOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.UserPreferenceOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.UserPreferenceOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTraces",
		reflect.TypeOf((*CfnMemoryTraces)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toXRay", GoMethod: "ToXRay"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMemoryTraces{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogs",
		reflect.TypeOf((*CfnRuntimeApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRuntimeApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeEndpointMixinProps",
		reflect.TypeOf((*CfnRuntimeEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeEndpointPropsMixin",
		reflect.TypeOf((*CfnRuntimeEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuntimeEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeLogsMixin",
		reflect.TypeOf((*CfnRuntimeLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuntimeLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeMixinProps",
		reflect.TypeOf((*CfnRuntimeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin",
		reflect.TypeOf((*CfnRuntimePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuntimePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.AgentRuntimeArtifactProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_AgentRuntimeArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.AuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_AuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.CodeConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.CodeProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.ContainerConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_ContainerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.CustomJWTAuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CustomJWTAuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.LifecycleConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_LifecycleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.RequestHeaderConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_RequestHeaderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.WorkloadIdentityDetailsProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_WorkloadIdentityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeTraces",
		reflect.TypeOf((*CfnRuntimeTraces)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toXRay", GoMethod: "ToXRay"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRuntimeTraces{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogs",
		reflect.TypeOf((*CfnRuntimeUsageLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRuntimeUsageLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogs",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnWorkloadIdentityApplicationLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityLogsMixin",
		reflect.TypeOf((*CfnWorkloadIdentityLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkloadIdentityLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityMixinProps",
		reflect.TypeOf((*CfnWorkloadIdentityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityPropsMixin",
		reflect.TypeOf((*CfnWorkloadIdentityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkloadIdentityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
