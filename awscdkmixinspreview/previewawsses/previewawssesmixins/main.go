package previewawssesmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetEventDestinationMixinProps",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetEventDestinationPropsMixin",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationSetEventDestinationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetEventDestinationPropsMixin.CloudWatchDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_CloudWatchDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetEventDestinationPropsMixin.DimensionConfigurationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_DimensionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetEventDestinationPropsMixin.EventBridgeDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_EventBridgeDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetEventDestinationPropsMixin.EventDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_EventDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetEventDestinationPropsMixin.KinesisFirehoseDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_KinesisFirehoseDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetEventDestinationPropsMixin.SnsDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_SnsDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetMixinProps",
		reflect.TypeOf((*CfnConfigurationSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.ArchivingOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_ArchivingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.ConditionThresholdProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_ConditionThresholdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.DashboardOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_DashboardOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.DeliveryOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_DeliveryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.GuardianOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_GuardianOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.OverallConfidenceThresholdProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_OverallConfidenceThresholdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.ReputationOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_ReputationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.SendingOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_SendingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.SuppressionOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_SuppressionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.TrackingOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_TrackingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.ValidationOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_ValidationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnConfigurationSetPropsMixin.VdmOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_VdmOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnContactListMixinProps",
		reflect.TypeOf((*CfnContactListMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnContactListPropsMixin",
		reflect.TypeOf((*CfnContactListPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContactListPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnContactListPropsMixin.TopicProperty",
		reflect.TypeOf((*CfnContactListPropsMixin_TopicProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnCustomVerificationEmailTemplateMixinProps",
		reflect.TypeOf((*CfnCustomVerificationEmailTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnCustomVerificationEmailTemplatePropsMixin",
		reflect.TypeOf((*CfnCustomVerificationEmailTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomVerificationEmailTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnDedicatedIpPoolMixinProps",
		reflect.TypeOf((*CfnDedicatedIpPoolMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnDedicatedIpPoolPropsMixin",
		reflect.TypeOf((*CfnDedicatedIpPoolPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDedicatedIpPoolPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityMixinProps",
		reflect.TypeOf((*CfnEmailIdentityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityPropsMixin",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEmailIdentityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityPropsMixin.ConfigurationSetAttributesProperty",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin_ConfigurationSetAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityPropsMixin.DkimAttributesProperty",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin_DkimAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityPropsMixin.DkimSigningAttributesProperty",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin_DkimSigningAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityPropsMixin.FeedbackAttributesProperty",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin_FeedbackAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityPropsMixin.MailFromAttributesProperty",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin_MailFromAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerAddonInstanceMixinProps",
		reflect.TypeOf((*CfnMailManagerAddonInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerAddonInstancePropsMixin",
		reflect.TypeOf((*CfnMailManagerAddonInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerAddonInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerAddonSubscriptionMixinProps",
		reflect.TypeOf((*CfnMailManagerAddonSubscriptionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerAddonSubscriptionPropsMixin",
		reflect.TypeOf((*CfnMailManagerAddonSubscriptionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerAddonSubscriptionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerAddressListMixinProps",
		reflect.TypeOf((*CfnMailManagerAddressListMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerAddressListPropsMixin",
		reflect.TypeOf((*CfnMailManagerAddressListPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerAddressListPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerArchiveMixinProps",
		reflect.TypeOf((*CfnMailManagerArchiveMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerArchivePropsMixin",
		reflect.TypeOf((*CfnMailManagerArchivePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerArchivePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerArchivePropsMixin.ArchiveRetentionProperty",
		reflect.TypeOf((*CfnMailManagerArchivePropsMixin_ArchiveRetentionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogs",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMailManagerIngressPointApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsDestProps",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnMailManagerIngressPointApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerIngressPointApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnMailManagerIngressPointApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnMailManagerIngressPointApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnMailManagerIngressPointApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnMailManagerIngressPointApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerIngressPointApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnMailManagerIngressPointApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnMailManagerIngressPointApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnMailManagerIngressPointApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsRecordFields",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"INGRESS_POINT_TYPE": CfnMailManagerIngressPointApplicationLogsRecordFields_INGRESS_POINT_TYPE,
			"INGRESS_POINT_ID": CfnMailManagerIngressPointApplicationLogsRecordFields_INGRESS_POINT_ID,
			"SENDER_IP_ADDRESS": CfnMailManagerIngressPointApplicationLogsRecordFields_SENDER_IP_ADDRESS,
			"SMTP_MAIL_FROM": CfnMailManagerIngressPointApplicationLogsRecordFields_SMTP_MAIL_FROM,
			"SMTP_HELO": CfnMailManagerIngressPointApplicationLogsRecordFields_SMTP_HELO,
			"TLS_PROTOCOL": CfnMailManagerIngressPointApplicationLogsRecordFields_TLS_PROTOCOL,
			"RECIPIENTS": CfnMailManagerIngressPointApplicationLogsRecordFields_RECIPIENTS,
			"INGRESS_POINT_METADATA": CfnMailManagerIngressPointApplicationLogsRecordFields_INGRESS_POINT_METADATA,
			"TLS_CIPHER_SUITE": CfnMailManagerIngressPointApplicationLogsRecordFields_TLS_CIPHER_SUITE,
			"MESSAGE_SIZE_BYTES": CfnMailManagerIngressPointApplicationLogsRecordFields_MESSAGE_SIZE_BYTES,
			"RULE_SET_ID": CfnMailManagerIngressPointApplicationLogsRecordFields_RULE_SET_ID,
			"RESOURCE_ARN": CfnMailManagerIngressPointApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnMailManagerIngressPointApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"MESSAGE_ID": CfnMailManagerIngressPointApplicationLogsRecordFields_MESSAGE_ID,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsS3Props",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointLogsMixin",
		reflect.TypeOf((*CfnMailManagerIngressPointLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerIngressPointLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointMixinProps",
		reflect.TypeOf((*CfnMailManagerIngressPointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointPropsMixin",
		reflect.TypeOf((*CfnMailManagerIngressPointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerIngressPointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointPropsMixin.IngressPointConfigurationProperty",
		reflect.TypeOf((*CfnMailManagerIngressPointPropsMixin_IngressPointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointPropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnMailManagerIngressPointPropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointPropsMixin.PrivateNetworkConfigurationProperty",
		reflect.TypeOf((*CfnMailManagerIngressPointPropsMixin_PrivateNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointPropsMixin.PublicNetworkConfigurationProperty",
		reflect.TypeOf((*CfnMailManagerIngressPointPropsMixin_PublicNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogs",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsDestProps",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsFirehoseProps",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsLogGroupProps",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat.S3",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_S3_JSON,
			"PLAIN": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_S3_PLAIN,
			"W3C": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_S3_W3C,
			"PARQUET": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"INGRESS_POINT_TYPE": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_INGRESS_POINT_TYPE,
			"INGRESS_POINT_ID": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_INGRESS_POINT_ID,
			"INGRESS_POINT_SESSION_ID": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_INGRESS_POINT_SESSION_ID,
			"TRAFFIC_POLICY_ID": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_TRAFFIC_POLICY_ID,
			"TRAFFIC_POLICY_EVALUATION": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_TRAFFIC_POLICY_EVALUATION,
			"TRAFFIC_POLICY_VERDICT": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_TRAFFIC_POLICY_VERDICT,
			"SENDER_IP_ADDRESS": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_SENDER_IP_ADDRESS,
			"SMTP_MAIL_FROM": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_SMTP_MAIL_FROM,
			"SMTP_HELO": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_SMTP_HELO,
			"TLS_PROTOCOL": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_TLS_PROTOCOL,
			"RECIPIENT": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_RECIPIENT,
			"TLS_CIPHER_SUITE": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_TLS_CIPHER_SUITE,
			"RESOURCE_ARN": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsS3Props",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRelayMixinProps",
		reflect.TypeOf((*CfnMailManagerRelayMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRelayPropsMixin",
		reflect.TypeOf((*CfnMailManagerRelayPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerRelayPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRelayPropsMixin.RelayAuthenticationProperty",
		reflect.TypeOf((*CfnMailManagerRelayPropsMixin_RelayAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogs",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMailManagerRuleSetApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsDestProps",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnMailManagerRuleSetApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerRuleSetApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnMailManagerRuleSetApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnMailManagerRuleSetApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnMailManagerRuleSetApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnMailManagerRuleSetApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerRuleSetApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnMailManagerRuleSetApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnMailManagerRuleSetApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnMailManagerRuleSetApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsRecordFields",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RULE_NAME": CfnMailManagerRuleSetApplicationLogsRecordFields_RULE_NAME,
			"RULE_INDEX": CfnMailManagerRuleSetApplicationLogsRecordFields_RULE_INDEX,
			"RECIPIENTS_MATCHED": CfnMailManagerRuleSetApplicationLogsRecordFields_RECIPIENTS_MATCHED,
			"ACTION_METADATA": CfnMailManagerRuleSetApplicationLogsRecordFields_ACTION_METADATA,
			"RESOURCE_ARN": CfnMailManagerRuleSetApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnMailManagerRuleSetApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"MESSAGE_ID": CfnMailManagerRuleSetApplicationLogsRecordFields_MESSAGE_ID,
			"RULE_SET_NAME": CfnMailManagerRuleSetApplicationLogsRecordFields_RULE_SET_NAME,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsS3Props",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetLogsMixin",
		reflect.TypeOf((*CfnMailManagerRuleSetLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerRuleSetLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetMixinProps",
		reflect.TypeOf((*CfnMailManagerRuleSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerRuleSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.AddHeaderActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_AddHeaderActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.AnalysisProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_AnalysisProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.ArchiveActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_ArchiveActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.DeliverToMailboxActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_DeliverToMailboxActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.DeliverToQBusinessActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_DeliverToQBusinessActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RelayActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RelayActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.ReplaceRecipientActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_ReplaceRecipientActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleBooleanExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleBooleanExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleBooleanToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleBooleanToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleConditionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleDmarcExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleDmarcExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleIpExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleIpExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleIpToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleIpToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleIsInAddressListProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleIsInAddressListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleNumberExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleNumberExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleNumberToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleNumberToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleStringExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleStringExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleStringToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleStringToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleVerdictExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleVerdictExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.RuleVerdictToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleVerdictToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.S3ActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_S3ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.SendActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_SendActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin.SnsActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_SnsActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyMixinProps",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerTrafficPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressAnalysisProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressAnalysisProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressBooleanExpressionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressBooleanExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressBooleanToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressBooleanToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressIpToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressIpToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressIpv4ExpressionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressIpv4ExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressIpv6ExpressionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressIpv6ExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressIpv6ToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressIpv6ToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressIsInAddressListProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressIsInAddressListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressStringExpressionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressStringExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressStringToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressStringToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressTlsProtocolExpressionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressTlsProtocolExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.IngressTlsProtocolToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressTlsProtocolToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.PolicyConditionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_PolicyConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerTrafficPolicyPropsMixin.PolicyStatementProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_PolicyStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMultiRegionEndpointMixinProps",
		reflect.TypeOf((*CfnMultiRegionEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMultiRegionEndpointPropsMixin",
		reflect.TypeOf((*CfnMultiRegionEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMultiRegionEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMultiRegionEndpointPropsMixin.DetailsProperty",
		reflect.TypeOf((*CfnMultiRegionEndpointPropsMixin_DetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMultiRegionEndpointPropsMixin.RouteDetailsItemsProperty",
		reflect.TypeOf((*CfnMultiRegionEndpointPropsMixin_RouteDetailsItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptFilterMixinProps",
		reflect.TypeOf((*CfnReceiptFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptFilterPropsMixin",
		reflect.TypeOf((*CfnReceiptFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReceiptFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptFilterPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnReceiptFilterPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptFilterPropsMixin.IpFilterProperty",
		reflect.TypeOf((*CfnReceiptFilterPropsMixin_IpFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRuleMixinProps",
		reflect.TypeOf((*CfnReceiptRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin",
		reflect.TypeOf((*CfnReceiptRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReceiptRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin.ActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin.AddHeaderActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_AddHeaderActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin.BounceActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_BounceActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin.ConnectActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_ConnectActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin.LambdaActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_LambdaActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin.RuleProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin.S3ActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_S3ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin.SNSActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_SNSActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin.StopActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_StopActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin.WorkmailActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_WorkmailActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRuleSetMixinProps",
		reflect.TypeOf((*CfnReceiptRuleSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRuleSetPropsMixin",
		reflect.TypeOf((*CfnReceiptRuleSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReceiptRuleSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnTemplateMixinProps",
		reflect.TypeOf((*CfnTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnTemplatePropsMixin",
		reflect.TypeOf((*CfnTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnTemplatePropsMixin.TemplateProperty",
		reflect.TypeOf((*CfnTemplatePropsMixin_TemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnTenantMixinProps",
		reflect.TypeOf((*CfnTenantMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnTenantPropsMixin",
		reflect.TypeOf((*CfnTenantPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTenantPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnTenantPropsMixin.ResourceAssociationProperty",
		reflect.TypeOf((*CfnTenantPropsMixin_ResourceAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnVdmAttributesMixinProps",
		reflect.TypeOf((*CfnVdmAttributesMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnVdmAttributesPropsMixin",
		reflect.TypeOf((*CfnVdmAttributesPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVdmAttributesPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnVdmAttributesPropsMixin.DashboardAttributesProperty",
		reflect.TypeOf((*CfnVdmAttributesPropsMixin_DashboardAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnVdmAttributesPropsMixin.GuardianAttributesProperty",
		reflect.TypeOf((*CfnVdmAttributesPropsMixin_GuardianAttributesProperty)(nil)).Elem(),
	)
}
