package previewawssammixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiMixinProps",
		reflect.TypeOf((*CfnApiMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiPropsMixin",
		reflect.TypeOf((*CfnApiPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApiPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiPropsMixin.AccessLogSettingProperty",
		reflect.TypeOf((*CfnApiPropsMixin_AccessLogSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiPropsMixin.AuthProperty",
		reflect.TypeOf((*CfnApiPropsMixin_AuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiPropsMixin.CanarySettingProperty",
		reflect.TypeOf((*CfnApiPropsMixin_CanarySettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiPropsMixin.CorsConfigurationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_CorsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiPropsMixin.DomainConfigurationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_DomainConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiPropsMixin.EndpointConfigurationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_EndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiPropsMixin.MutualTlsAuthenticationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_MutualTlsAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiPropsMixin.Route53ConfigurationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_Route53ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApiPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnApplicationPropsMixin.ApplicationLocationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionMixinProps",
		reflect.TypeOf((*CfnFunctionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin",
		reflect.TypeOf((*CfnFunctionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFunctionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.AlexaSkillEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_AlexaSkillEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.ApiEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ApiEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.AuthProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_AuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.AuthResourcePolicyProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_AuthResourcePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.BucketSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_BucketSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.CloudWatchEventEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CloudWatchEventEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.CloudWatchLogsEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CloudWatchLogsEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.CognitoEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CognitoEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.CollectionSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CollectionSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.CorsConfigurationProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CorsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.DeadLetterQueueProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DeadLetterQueueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.DeploymentPreferenceProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DeploymentPreferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.DestinationConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.DomainSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DomainSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.DynamoDBEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DynamoDBEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.EmptySAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EmptySAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.EphemeralStorageProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EphemeralStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.EventBridgeRuleEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EventBridgeRuleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.EventInvokeConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EventInvokeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.EventInvokeDestinationConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EventInvokeDestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.EventSourceProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.FileSystemConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.FunctionEnvironmentProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionEnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.FunctionSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.FunctionUrlConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionUrlConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.HooksProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_HooksProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.HttpApiEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_HttpApiEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.HttpApiFunctionAuthProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_HttpApiFunctionAuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.IAMPolicyDocumentProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_IAMPolicyDocumentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.IdentitySAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_IdentitySAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.ImageConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.IoTRuleEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_IoTRuleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.KeySAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_KeySAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.KinesisEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_KinesisEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.LogGroupSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_LogGroupSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.ParameterNameSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ParameterNameSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.ProvisionedConcurrencyConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ProvisionedConcurrencyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.QueueSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_QueueSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.RequestModelProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_RequestModelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.RequestParameterProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_RequestParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.RouteSettingsProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_RouteSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.S3EventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_S3EventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.S3KeyFilterProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_S3KeyFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.S3KeyFilterRuleProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_S3KeyFilterRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.S3NotificationFilterProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_S3NotificationFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.SAMPolicyTemplateProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SAMPolicyTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.SNSEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SNSEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.SQSEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SQSEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.ScheduleEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ScheduleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.SecretArnSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SecretArnSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.StateMachineSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_StateMachineSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.StreamSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_StreamSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.TableSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_TableSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.TableStreamSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_TableStreamSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.TopicSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_TopicSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiMixinProps",
		reflect.TypeOf((*CfnHttpApiMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin",
		reflect.TypeOf((*CfnHttpApiPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHttpApiPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin.AccessLogSettingProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_AccessLogSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin.CorsConfigurationObjectProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_CorsConfigurationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin.HttpApiAuthProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_HttpApiAuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin.HttpApiDomainConfigurationProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_HttpApiDomainConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin.MutualTlsAuthenticationProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_MutualTlsAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin.Route53ConfigurationProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_Route53ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin.RouteSettingsProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_RouteSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnHttpApiPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnLayerVersionMixinProps",
		reflect.TypeOf((*CfnLayerVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnLayerVersionPropsMixin",
		reflect.TypeOf((*CfnLayerVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLayerVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnLayerVersionPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnLayerVersionPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnSimpleTableMixinProps",
		reflect.TypeOf((*CfnSimpleTableMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnSimpleTablePropsMixin",
		reflect.TypeOf((*CfnSimpleTablePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSimpleTablePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnSimpleTablePropsMixin.PrimaryKeyProperty",
		reflect.TypeOf((*CfnSimpleTablePropsMixin_PrimaryKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnSimpleTablePropsMixin.ProvisionedThroughputProperty",
		reflect.TypeOf((*CfnSimpleTablePropsMixin_ProvisionedThroughputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnSimpleTablePropsMixin.SSESpecificationProperty",
		reflect.TypeOf((*CfnSimpleTablePropsMixin_SSESpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachineMixinProps",
		reflect.TypeOf((*CfnStateMachineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin",
		reflect.TypeOf((*CfnStateMachinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateMachinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.ApiEventProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_ApiEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.CloudWatchEventEventProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_CloudWatchEventEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.CloudWatchLogsLogGroupProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_CloudWatchLogsLogGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.EventBridgeRuleEventProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_EventBridgeRuleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.EventSourceProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_EventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.FunctionSAMPTProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_FunctionSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.IAMPolicyDocumentProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_IAMPolicyDocumentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.LogDestinationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_LogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.SAMPolicyTemplateProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_SAMPolicyTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.ScheduleEventProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_ScheduleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.StateMachineSAMPTProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_StateMachineSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnStateMachinePropsMixin.TracingConfigurationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_TracingConfigurationProperty)(nil)).Elem(),
	)
}
