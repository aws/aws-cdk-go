package awsmsk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.BrokerLogging",
		reflect.TypeOf((*BrokerLogging)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_msk.CfnBatchScramSecret",
		reflect.TypeOf((*CfnBatchScramSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "secretArnList", GoGetter: "SecretArnList"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBatchScramSecret{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnBatchScramSecretProps",
		reflect.TypeOf((*CfnBatchScramSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_msk.CfnCluster",
		reflect.TypeOf((*CfnCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "brokerNodeGroupInfo", GoGetter: "BrokerNodeGroupInfo"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientAuthentication", GoGetter: "ClientAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "configurationInfo", GoGetter: "ConfigurationInfo"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "currentVersion", GoGetter: "CurrentVersion"},
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
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "openMonitoring", GoGetter: "OpenMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storageMode", GoGetter: "StorageMode"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
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
		"monocdk.aws_msk.CfnCluster.BrokerLogsProperty",
		reflect.TypeOf((*CfnCluster_BrokerLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.BrokerNodeGroupInfoProperty",
		reflect.TypeOf((*CfnCluster_BrokerNodeGroupInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.ClientAuthenticationProperty",
		reflect.TypeOf((*CfnCluster_ClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.CloudWatchLogsProperty",
		reflect.TypeOf((*CfnCluster_CloudWatchLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.ConfigurationInfoProperty",
		reflect.TypeOf((*CfnCluster_ConfigurationInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.ConnectivityInfoProperty",
		reflect.TypeOf((*CfnCluster_ConnectivityInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.EBSStorageInfoProperty",
		reflect.TypeOf((*CfnCluster_EBSStorageInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.EncryptionAtRestProperty",
		reflect.TypeOf((*CfnCluster_EncryptionAtRestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.EncryptionInTransitProperty",
		reflect.TypeOf((*CfnCluster_EncryptionInTransitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.EncryptionInfoProperty",
		reflect.TypeOf((*CfnCluster_EncryptionInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.FirehoseProperty",
		reflect.TypeOf((*CfnCluster_FirehoseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.IamProperty",
		reflect.TypeOf((*CfnCluster_IamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.JmxExporterProperty",
		reflect.TypeOf((*CfnCluster_JmxExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.LoggingInfoProperty",
		reflect.TypeOf((*CfnCluster_LoggingInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.NodeExporterProperty",
		reflect.TypeOf((*CfnCluster_NodeExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.OpenMonitoringProperty",
		reflect.TypeOf((*CfnCluster_OpenMonitoringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.PrometheusProperty",
		reflect.TypeOf((*CfnCluster_PrometheusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.ProvisionedThroughputProperty",
		reflect.TypeOf((*CfnCluster_ProvisionedThroughputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.PublicAccessProperty",
		reflect.TypeOf((*CfnCluster_PublicAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.S3Property",
		reflect.TypeOf((*CfnCluster_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.SaslProperty",
		reflect.TypeOf((*CfnCluster_SaslProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.ScramProperty",
		reflect.TypeOf((*CfnCluster_ScramProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.StorageInfoProperty",
		reflect.TypeOf((*CfnCluster_StorageInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.TlsProperty",
		reflect.TypeOf((*CfnCluster_TlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.UnauthenticatedProperty",
		reflect.TypeOf((*CfnCluster_UnauthenticatedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.VpcConnectivityClientAuthenticationProperty",
		reflect.TypeOf((*CfnCluster_VpcConnectivityClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.VpcConnectivityIamProperty",
		reflect.TypeOf((*CfnCluster_VpcConnectivityIamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.VpcConnectivityProperty",
		reflect.TypeOf((*CfnCluster_VpcConnectivityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.VpcConnectivitySaslProperty",
		reflect.TypeOf((*CfnCluster_VpcConnectivitySaslProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.VpcConnectivityScramProperty",
		reflect.TypeOf((*CfnCluster_VpcConnectivityScramProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnCluster.VpcConnectivityTlsProperty",
		reflect.TypeOf((*CfnCluster_VpcConnectivityTlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_msk.CfnClusterPolicy",
		reflect.TypeOf((*CfnClusterPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCurrentVersion", GoGetter: "AttrCurrentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnClusterPolicyProps",
		reflect.TypeOf((*CfnClusterPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnClusterProps",
		reflect.TypeOf((*CfnClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_msk.CfnConfiguration",
		reflect.TypeOf((*CfnConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kafkaVersionsList", GoGetter: "KafkaVersionsList"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "serverProperties", GoGetter: "ServerProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnConfigurationProps",
		reflect.TypeOf((*CfnConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_msk.CfnServerlessCluster",
		reflect.TypeOf((*CfnServerlessCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientAuthentication", GoGetter: "ClientAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfigs", GoGetter: "VpcConfigs"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServerlessCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnServerlessCluster.ClientAuthenticationProperty",
		reflect.TypeOf((*CfnServerlessCluster_ClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnServerlessCluster.IamProperty",
		reflect.TypeOf((*CfnServerlessCluster_IamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnServerlessCluster.SaslProperty",
		reflect.TypeOf((*CfnServerlessCluster_SaslProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnServerlessCluster.VpcConfigProperty",
		reflect.TypeOf((*CfnServerlessCluster_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnServerlessClusterProps",
		reflect.TypeOf((*CfnServerlessClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_msk.CfnVpcConnection",
		reflect.TypeOf((*CfnVpcConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientSubnets", GoGetter: "ClientSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "targetClusterArn", GoGetter: "TargetClusterArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcConnection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.CfnVpcConnectionProps",
		reflect.TypeOf((*CfnVpcConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_msk.ClientAuthentication",
		reflect.TypeOf((*ClientAuthentication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "saslProps", GoGetter: "SaslProps"},
			_jsii_.MemberProperty{JsiiProperty: "tlsProps", GoGetter: "TlsProps"},
		},
		func() interface{} {
			return &jsiiProxy_ClientAuthentication{}
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_msk.ClientBrokerEncryption",
		reflect.TypeOf((*ClientBrokerEncryption)(nil)).Elem(),
		map[string]interface{}{
			"TLS": ClientBrokerEncryption_TLS,
			"TLS_PLAINTEXT": ClientBrokerEncryption_TLS_PLAINTEXT,
			"PLAINTEXT": ClientBrokerEncryption_PLAINTEXT,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_msk.Cluster",
		reflect.TypeOf((*Cluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addUser", GoMethod: "AddUser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapBrokers", GoGetter: "BootstrapBrokers"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapBrokersSaslScram", GoGetter: "BootstrapBrokersSaslScram"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapBrokersTls", GoGetter: "BootstrapBrokersTls"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "saslScramAuthenticationKey", GoGetter: "SaslScramAuthenticationKey"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "zookeeperConnectionString", GoGetter: "ZookeeperConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "zookeeperConnectionStringTls", GoGetter: "ZookeeperConnectionStringTls"},
		},
		func() interface{} {
			j := jsiiProxy_Cluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICluster)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.ClusterConfigurationInfo",
		reflect.TypeOf((*ClusterConfigurationInfo)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_msk.ClusterMonitoringLevel",
		reflect.TypeOf((*ClusterMonitoringLevel)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": ClusterMonitoringLevel_DEFAULT,
			"PER_BROKER": ClusterMonitoringLevel_PER_BROKER,
			"PER_TOPIC_PER_BROKER": ClusterMonitoringLevel_PER_TOPIC_PER_BROKER,
			"PER_TOPIC_PER_PARTITION": ClusterMonitoringLevel_PER_TOPIC_PER_PARTITION,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.ClusterProps",
		reflect.TypeOf((*ClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.EbsStorageInfo",
		reflect.TypeOf((*EbsStorageInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.EncryptionInTransitConfig",
		reflect.TypeOf((*EncryptionInTransitConfig)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_msk.ICluster",
		reflect.TypeOf((*ICluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ICluster{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_msk.KafkaVersion",
		reflect.TypeOf((*KafkaVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_KafkaVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.MonitoringConfiguration",
		reflect.TypeOf((*MonitoringConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.S3LoggingConfiguration",
		reflect.TypeOf((*S3LoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.SaslAuthProps",
		reflect.TypeOf((*SaslAuthProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_msk.TlsAuthProps",
		reflect.TypeOf((*TlsAuthProps)(nil)).Elem(),
	)
}
