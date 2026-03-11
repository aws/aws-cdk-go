package awssam

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiMixinProps",
		reflect.TypeOf((*CfnApiMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiPropsMixin",
		reflect.TypeOf((*CfnApiPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApiPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiPropsMixin.AccessLogSettingProperty",
		reflect.TypeOf((*CfnApiPropsMixin_AccessLogSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiPropsMixin.AuthProperty",
		reflect.TypeOf((*CfnApiPropsMixin_AuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiPropsMixin.CanarySettingProperty",
		reflect.TypeOf((*CfnApiPropsMixin_CanarySettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiPropsMixin.CorsConfigurationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_CorsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiPropsMixin.DomainConfigurationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_DomainConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiPropsMixin.EndpointConfigurationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_EndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiPropsMixin.MutualTlsAuthenticationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_MutualTlsAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiPropsMixin.Route53ConfigurationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_Route53ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApiPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnApiPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnApplicationPropsMixin.ApplicationLocationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionMixinProps",
		reflect.TypeOf((*CfnFunctionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin",
		reflect.TypeOf((*CfnFunctionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFunctionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.AlexaSkillEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_AlexaSkillEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.ApiEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ApiEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.AuthProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_AuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.AuthResourcePolicyProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_AuthResourcePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.BucketSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_BucketSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.CloudWatchEventEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CloudWatchEventEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.CloudWatchLogsEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CloudWatchLogsEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.CognitoEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CognitoEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.CollectionSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CollectionSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.CorsConfigurationProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_CorsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.DeadLetterQueueProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DeadLetterQueueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.DeploymentPreferenceProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DeploymentPreferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.DestinationConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.DomainSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DomainSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.DynamoDBEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_DynamoDBEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.EmptySAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EmptySAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.EphemeralStorageProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EphemeralStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.EventBridgeRuleEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EventBridgeRuleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.EventInvokeConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EventInvokeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.EventInvokeDestinationConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EventInvokeDestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.EventSourceProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_EventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.FileSystemConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.FunctionEnvironmentProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionEnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.FunctionSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.FunctionUrlConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionUrlConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.HooksProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_HooksProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.HttpApiEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_HttpApiEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.HttpApiFunctionAuthProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_HttpApiFunctionAuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.IAMPolicyDocumentProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_IAMPolicyDocumentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.IdentitySAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_IdentitySAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.ImageConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ImageConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.IoTRuleEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_IoTRuleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.KeySAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_KeySAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.KinesisEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_KinesisEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.LogGroupSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_LogGroupSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.ParameterNameSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ParameterNameSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.ProvisionedConcurrencyConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ProvisionedConcurrencyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.QueueSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_QueueSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.RequestModelProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_RequestModelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.RequestParameterProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_RequestParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.RouteSettingsProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_RouteSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.S3EventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_S3EventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.S3KeyFilterProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_S3KeyFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.S3KeyFilterRuleProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_S3KeyFilterRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.S3NotificationFilterProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_S3NotificationFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.SAMPolicyTemplateProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SAMPolicyTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.SNSEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SNSEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.SQSEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SQSEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.ScheduleEventProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_ScheduleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.SecretArnSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_SecretArnSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.StateMachineSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_StateMachineSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.StreamSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_StreamSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.TableSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_TableSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.TableStreamSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_TableStreamSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.TopicSAMPTProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_TopicSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnFunctionPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnHttpApiMixinProps",
		reflect.TypeOf((*CfnHttpApiMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnHttpApiPropsMixin",
		reflect.TypeOf((*CfnHttpApiPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHttpApiPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnHttpApiPropsMixin.AccessLogSettingProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_AccessLogSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnHttpApiPropsMixin.CorsConfigurationObjectProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_CorsConfigurationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnHttpApiPropsMixin.HttpApiAuthProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_HttpApiAuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnHttpApiPropsMixin.HttpApiDomainConfigurationProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_HttpApiDomainConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnHttpApiPropsMixin.MutualTlsAuthenticationProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_MutualTlsAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnHttpApiPropsMixin.Route53ConfigurationProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_Route53ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnHttpApiPropsMixin.RouteSettingsProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_RouteSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnHttpApiPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnHttpApiPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnLayerVersionMixinProps",
		reflect.TypeOf((*CfnLayerVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnLayerVersionPropsMixin",
		reflect.TypeOf((*CfnLayerVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLayerVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnLayerVersionPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnLayerVersionPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnSimpleTableMixinProps",
		reflect.TypeOf((*CfnSimpleTableMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnSimpleTablePropsMixin",
		reflect.TypeOf((*CfnSimpleTablePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSimpleTablePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnSimpleTablePropsMixin.PrimaryKeyProperty",
		reflect.TypeOf((*CfnSimpleTablePropsMixin_PrimaryKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnSimpleTablePropsMixin.ProvisionedThroughputProperty",
		reflect.TypeOf((*CfnSimpleTablePropsMixin_ProvisionedThroughputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnSimpleTablePropsMixin.SSESpecificationProperty",
		reflect.TypeOf((*CfnSimpleTablePropsMixin_SSESpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachineMixinProps",
		reflect.TypeOf((*CfnStateMachineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin",
		reflect.TypeOf((*CfnStateMachinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateMachinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.ApiEventProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_ApiEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.CloudWatchEventEventProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_CloudWatchEventEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.CloudWatchLogsLogGroupProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_CloudWatchLogsLogGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.EventBridgeRuleEventProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_EventBridgeRuleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.EventSourceProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_EventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.FunctionSAMPTProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_FunctionSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.IAMPolicyDocumentProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_IAMPolicyDocumentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.LogDestinationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_LogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.SAMPolicyTemplateProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_SAMPolicyTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.ScheduleEventProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_ScheduleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.StateMachineSAMPTProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_StateMachineSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_sam.CfnStateMachinePropsMixin.TracingConfigurationProperty",
		reflect.TypeOf((*CfnStateMachinePropsMixin_TracingConfigurationProperty)(nil)).Elem(),
	)
}
