package previewawsmskmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnBatchScramSecretMixinProps",
		reflect.TypeOf((*CfnBatchScramSecretMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnBatchScramSecretPropsMixin",
		reflect.TypeOf((*CfnBatchScramSecretPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBatchScramSecretPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogs",
		reflect.TypeOf((*CfnClusterBrokerLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterBrokerLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterLogsMixin",
		reflect.TypeOf((*CfnClusterLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPolicyMixinProps",
		reflect.TypeOf((*CfnClusterPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPolicyPropsMixin",
		reflect.TypeOf((*CfnClusterPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin",
		reflect.TypeOf((*CfnClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.BrokerLogsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_BrokerLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.BrokerNodeGroupInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_BrokerNodeGroupInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.ClientAuthenticationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.CloudWatchLogsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_CloudWatchLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.ConfigurationInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ConfigurationInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.ConnectivityInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ConnectivityInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.EBSStorageInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EBSStorageInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.EncryptionAtRestProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EncryptionAtRestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.EncryptionInTransitProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EncryptionInTransitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.EncryptionInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EncryptionInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.FirehoseProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_FirehoseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.IamProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_IamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.JmxExporterProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_JmxExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.LoggingInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_LoggingInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.NodeExporterProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_NodeExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.OpenMonitoringProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OpenMonitoringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.PrometheusProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_PrometheusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.ProvisionedThroughputProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ProvisionedThroughputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.PublicAccessProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_PublicAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.RebalancingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_RebalancingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.S3Property",
		reflect.TypeOf((*CfnClusterPropsMixin_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.SaslProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SaslProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.ScramProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScramProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.StorageInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_StorageInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.TlsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_TlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.UnauthenticatedProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_UnauthenticatedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.VpcConnectivityClientAuthenticationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivityClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.VpcConnectivityIamProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivityIamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.VpcConnectivityProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.VpcConnectivitySaslProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivitySaslProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.VpcConnectivityScramProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivityScramProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterPropsMixin.VpcConnectivityTlsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivityTlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnConfigurationMixinProps",
		reflect.TypeOf((*CfnConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnConfigurationPropsMixin",
		reflect.TypeOf((*CfnConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnConfigurationPropsMixin.LatestRevisionProperty",
		reflect.TypeOf((*CfnConfigurationPropsMixin_LatestRevisionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorMixinProps",
		reflect.TypeOf((*CfnReplicatorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin",
		reflect.TypeOf((*CfnReplicatorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReplicatorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin.AmazonMskClusterProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_AmazonMskClusterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin.ConsumerGroupReplicationProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_ConsumerGroupReplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin.KafkaClusterClientVpcConfigProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_KafkaClusterClientVpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin.KafkaClusterProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_KafkaClusterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin.ReplicationInfoProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_ReplicationInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin.ReplicationStartingPositionProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_ReplicationStartingPositionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin.ReplicationTopicNameConfigurationProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_ReplicationTopicNameConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorPropsMixin.TopicReplicationProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_TopicReplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnServerlessClusterMixinProps",
		reflect.TypeOf((*CfnServerlessClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnServerlessClusterPropsMixin",
		reflect.TypeOf((*CfnServerlessClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServerlessClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnServerlessClusterPropsMixin.ClientAuthenticationProperty",
		reflect.TypeOf((*CfnServerlessClusterPropsMixin_ClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnServerlessClusterPropsMixin.IamProperty",
		reflect.TypeOf((*CfnServerlessClusterPropsMixin_IamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnServerlessClusterPropsMixin.SaslProperty",
		reflect.TypeOf((*CfnServerlessClusterPropsMixin_SaslProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnServerlessClusterPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnServerlessClusterPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnVpcConnectionMixinProps",
		reflect.TypeOf((*CfnVpcConnectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnVpcConnectionPropsMixin",
		reflect.TypeOf((*CfnVpcConnectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcConnectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
