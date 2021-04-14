package awsmsk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_msk.CfnCluster",
		reflect.TypeOf((*CfnCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "brokerNodeGroupInfo", GoGetter: "BrokerNodeGroupInfo"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientAuthentication", GoGetter: "ClientAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "configurationInfo", GoGetter: "ConfigurationInfo"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionInfo", GoGetter: "EncryptionInfo"},
			_jsii_.MemberProperty{JsiiProperty: "enhancedMonitoring", GoGetter: "EnhancedMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kafkaVersion", GoGetter: "KafkaVersion"},
			_jsii_.MemberProperty{JsiiProperty: "loggingInfo", GoGetter: "LoggingInfo"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfBrokerNodes", GoGetter: "NumberOfBrokerNodes"},
			_jsii_.MemberProperty{JsiiProperty: "openMonitoring", GoGetter: "OpenMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.BrokerLogsProperty",
		reflect.TypeOf((*CfnCluster_BrokerLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.BrokerNodeGroupInfoProperty",
		reflect.TypeOf((*CfnCluster_BrokerNodeGroupInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.ClientAuthenticationProperty",
		reflect.TypeOf((*CfnCluster_ClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.CloudWatchLogsProperty",
		reflect.TypeOf((*CfnCluster_CloudWatchLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.ConfigurationInfoProperty",
		reflect.TypeOf((*CfnCluster_ConfigurationInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.EBSStorageInfoProperty",
		reflect.TypeOf((*CfnCluster_EBSStorageInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.EncryptionAtRestProperty",
		reflect.TypeOf((*CfnCluster_EncryptionAtRestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.EncryptionInTransitProperty",
		reflect.TypeOf((*CfnCluster_EncryptionInTransitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.EncryptionInfoProperty",
		reflect.TypeOf((*CfnCluster_EncryptionInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.FirehoseProperty",
		reflect.TypeOf((*CfnCluster_FirehoseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.JmxExporterProperty",
		reflect.TypeOf((*CfnCluster_JmxExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.LoggingInfoProperty",
		reflect.TypeOf((*CfnCluster_LoggingInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.NodeExporterProperty",
		reflect.TypeOf((*CfnCluster_NodeExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.OpenMonitoringProperty",
		reflect.TypeOf((*CfnCluster_OpenMonitoringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.PrometheusProperty",
		reflect.TypeOf((*CfnCluster_PrometheusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.S3Property",
		reflect.TypeOf((*CfnCluster_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.SaslProperty",
		reflect.TypeOf((*CfnCluster_SaslProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.ScramProperty",
		reflect.TypeOf((*CfnCluster_ScramProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.StorageInfoProperty",
		reflect.TypeOf((*CfnCluster_StorageInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnCluster.TlsProperty",
		reflect.TypeOf((*CfnCluster_TlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_msk.CfnClusterProps",
		reflect.TypeOf((*CfnClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_msk.Cluster",
		reflect.TypeOf((*Cluster)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Cluster{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_msk.ICluster",
		reflect.TypeOf((*ICluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ICluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
}
