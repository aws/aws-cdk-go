package awsmsk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnBatchScramSecretMixinProps",
		reflect.TypeOf((*CfnBatchScramSecretMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnBatchScramSecretPropsMixin",
		reflect.TypeOf((*CfnBatchScramSecretPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBatchScramSecretPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPolicyMixinProps",
		reflect.TypeOf((*CfnClusterPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPolicyPropsMixin",
		reflect.TypeOf((*CfnClusterPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin",
		reflect.TypeOf((*CfnClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.BrokerLogsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_BrokerLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.BrokerNodeGroupInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_BrokerNodeGroupInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.ClientAuthenticationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.CloudWatchLogsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_CloudWatchLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.ConfigurationInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ConfigurationInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.ConnectivityInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ConnectivityInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.EBSStorageInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EBSStorageInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.EncryptionAtRestProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EncryptionAtRestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.EncryptionInTransitProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EncryptionInTransitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.EncryptionInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EncryptionInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.FirehoseProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_FirehoseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.IamProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_IamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.JmxExporterProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_JmxExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.LoggingInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_LoggingInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.NodeExporterProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_NodeExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.OpenMonitoringProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OpenMonitoringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.PrometheusProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_PrometheusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.ProvisionedThroughputProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ProvisionedThroughputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.PublicAccessProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_PublicAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.RebalancingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_RebalancingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.S3Property",
		reflect.TypeOf((*CfnClusterPropsMixin_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.SaslProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SaslProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.ScramProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScramProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.StorageInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_StorageInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.TlsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_TlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.UnauthenticatedProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_UnauthenticatedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.VpcConnectivityClientAuthenticationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivityClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.VpcConnectivityIamProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivityIamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.VpcConnectivityProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.VpcConnectivitySaslProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivitySaslProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.VpcConnectivityScramProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivityScramProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnClusterPropsMixin.VpcConnectivityTlsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VpcConnectivityTlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnConfigurationMixinProps",
		reflect.TypeOf((*CfnConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnConfigurationPropsMixin",
		reflect.TypeOf((*CfnConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnConfigurationPropsMixin.LatestRevisionProperty",
		reflect.TypeOf((*CfnConfigurationPropsMixin_LatestRevisionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnReplicatorMixinProps",
		reflect.TypeOf((*CfnReplicatorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnReplicatorPropsMixin",
		reflect.TypeOf((*CfnReplicatorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReplicatorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnReplicatorPropsMixin.AmazonMskClusterProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_AmazonMskClusterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnReplicatorPropsMixin.ConsumerGroupReplicationProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_ConsumerGroupReplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnReplicatorPropsMixin.KafkaClusterClientVpcConfigProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_KafkaClusterClientVpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnReplicatorPropsMixin.KafkaClusterProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_KafkaClusterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnReplicatorPropsMixin.ReplicationInfoProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_ReplicationInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnReplicatorPropsMixin.ReplicationStartingPositionProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_ReplicationStartingPositionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnReplicatorPropsMixin.ReplicationTopicNameConfigurationProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_ReplicationTopicNameConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnReplicatorPropsMixin.TopicReplicationProperty",
		reflect.TypeOf((*CfnReplicatorPropsMixin_TopicReplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnServerlessClusterMixinProps",
		reflect.TypeOf((*CfnServerlessClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnServerlessClusterPropsMixin",
		reflect.TypeOf((*CfnServerlessClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServerlessClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnServerlessClusterPropsMixin.ClientAuthenticationProperty",
		reflect.TypeOf((*CfnServerlessClusterPropsMixin_ClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnServerlessClusterPropsMixin.IamProperty",
		reflect.TypeOf((*CfnServerlessClusterPropsMixin_IamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnServerlessClusterPropsMixin.SaslProperty",
		reflect.TypeOf((*CfnServerlessClusterPropsMixin_SaslProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnServerlessClusterPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnServerlessClusterPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnTopicMixinProps",
		reflect.TypeOf((*CfnTopicMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnTopicPropsMixin",
		reflect.TypeOf((*CfnTopicPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTopicPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnVpcConnectionMixinProps",
		reflect.TypeOf((*CfnVpcConnectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_msk.CfnVpcConnectionPropsMixin",
		reflect.TypeOf((*CfnVpcConnectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcConnectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
