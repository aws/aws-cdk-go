package awskafkaconnect

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		reflect.TypeOf((*CfnConnector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrConnectorArn", GoGetter: "AttrConnectorArn"},
			_jsii_.MemberProperty{JsiiProperty: "capacity", GoGetter: "Capacity"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectorConfiguration", GoGetter: "ConnectorConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "connectorDescription", GoGetter: "ConnectorDescription"},
			_jsii_.MemberProperty{JsiiProperty: "connectorName", GoGetter: "ConnectorName"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kafkaCluster", GoGetter: "KafkaCluster"},
			_jsii_.MemberProperty{JsiiProperty: "kafkaClusterClientAuthentication", GoGetter: "KafkaClusterClientAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "kafkaClusterEncryptionInTransit", GoGetter: "KafkaClusterEncryptionInTransit"},
			_jsii_.MemberProperty{JsiiProperty: "kafkaConnectVersion", GoGetter: "KafkaConnectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "plugins", GoGetter: "Plugins"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "serviceExecutionRoleArn", GoGetter: "ServiceExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "workerConfiguration", GoGetter: "WorkerConfiguration"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnector{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.ApacheKafkaClusterProperty",
		reflect.TypeOf((*CfnConnector_ApacheKafkaClusterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.AutoScalingProperty",
		reflect.TypeOf((*CfnConnector_AutoScalingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.CapacityProperty",
		reflect.TypeOf((*CfnConnector_CapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.CloudWatchLogsLogDeliveryProperty",
		reflect.TypeOf((*CfnConnector_CloudWatchLogsLogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.CustomPluginProperty",
		reflect.TypeOf((*CfnConnector_CustomPluginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.FirehoseLogDeliveryProperty",
		reflect.TypeOf((*CfnConnector_FirehoseLogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.KafkaClusterClientAuthenticationProperty",
		reflect.TypeOf((*CfnConnector_KafkaClusterClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.KafkaClusterEncryptionInTransitProperty",
		reflect.TypeOf((*CfnConnector_KafkaClusterEncryptionInTransitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.KafkaClusterProperty",
		reflect.TypeOf((*CfnConnector_KafkaClusterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.LogDeliveryProperty",
		reflect.TypeOf((*CfnConnector_LogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.PluginProperty",
		reflect.TypeOf((*CfnConnector_PluginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.ProvisionedCapacityProperty",
		reflect.TypeOf((*CfnConnector_ProvisionedCapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.S3LogDeliveryProperty",
		reflect.TypeOf((*CfnConnector_S3LogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.ScaleInPolicyProperty",
		reflect.TypeOf((*CfnConnector_ScaleInPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.ScaleOutPolicyProperty",
		reflect.TypeOf((*CfnConnector_ScaleOutPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.VpcProperty",
		reflect.TypeOf((*CfnConnector_VpcProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.WorkerConfigurationProperty",
		reflect.TypeOf((*CfnConnector_WorkerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector.WorkerLogDeliveryProperty",
		reflect.TypeOf((*CfnConnector_WorkerLogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnectorProps",
		reflect.TypeOf((*CfnConnectorProps)(nil)).Elem(),
	)
}
