// The CDK Construct Library for AWS::MSK
package awscdkmskalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-msk-alpha.BrokerLogging",
		reflect.TypeOf((*BrokerLogging)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-msk-alpha.ClientAuthentication",
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
		"@aws-cdk/aws-msk-alpha.ClientBrokerEncryption",
		reflect.TypeOf((*ClientBrokerEncryption)(nil)).Elem(),
		map[string]interface{}{
			"TLS": ClientBrokerEncryption_TLS,
			"TLS_PLAINTEXT": ClientBrokerEncryption_TLS_PLAINTEXT,
			"PLAINTEXT": ClientBrokerEncryption_PLAINTEXT,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-msk-alpha.Cluster",
		reflect.TypeOf((*Cluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addUser", GoMethod: "AddUser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapBrokers", GoGetter: "BootstrapBrokers"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapBrokersSaslIam", GoGetter: "BootstrapBrokersSaslIam"},
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
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "saslScramAuthenticationKey", GoGetter: "SaslScramAuthenticationKey"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
		"@aws-cdk/aws-msk-alpha.ClusterConfigurationInfo",
		reflect.TypeOf((*ClusterConfigurationInfo)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-msk-alpha.ClusterMonitoringLevel",
		reflect.TypeOf((*ClusterMonitoringLevel)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": ClusterMonitoringLevel_DEFAULT,
			"PER_BROKER": ClusterMonitoringLevel_PER_BROKER,
			"PER_TOPIC_PER_BROKER": ClusterMonitoringLevel_PER_TOPIC_PER_BROKER,
			"PER_TOPIC_PER_PARTITION": ClusterMonitoringLevel_PER_TOPIC_PER_PARTITION,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-msk-alpha.ClusterProps",
		reflect.TypeOf((*ClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-msk-alpha.EbsStorageInfo",
		reflect.TypeOf((*EbsStorageInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-msk-alpha.EncryptionInTransitConfig",
		reflect.TypeOf((*EncryptionInTransitConfig)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-msk-alpha.ICluster",
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
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		reflect.TypeOf((*KafkaVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "features", GoGetter: "Features"},
			_jsii_.MemberMethod{JsiiMethod: "isTieredStorageCompatible", GoMethod: "IsTieredStorageCompatible"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_KafkaVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-msk-alpha.KafkaVersionFeatures",
		reflect.TypeOf((*KafkaVersionFeatures)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-msk-alpha.MonitoringConfiguration",
		reflect.TypeOf((*MonitoringConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-msk-alpha.S3LoggingConfiguration",
		reflect.TypeOf((*S3LoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-msk-alpha.SaslAuthProps",
		reflect.TypeOf((*SaslAuthProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-msk-alpha.SaslTlsAuthProps",
		reflect.TypeOf((*SaslTlsAuthProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-msk-alpha.StorageMode",
		reflect.TypeOf((*StorageMode)(nil)).Elem(),
		map[string]interface{}{
			"LOCAL": StorageMode_LOCAL,
			"TIERED": StorageMode_TIERED,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-msk-alpha.TlsAuthProps",
		reflect.TypeOf((*TlsAuthProps)(nil)).Elem(),
	)
}
