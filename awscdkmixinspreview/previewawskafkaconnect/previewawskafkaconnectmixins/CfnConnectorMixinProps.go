package previewawskafkaconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorMixinProps := &CfnConnectorMixinProps{
//   	Capacity: &CapacityProperty{
//   		AutoScaling: &AutoScalingProperty{
//   			MaxWorkerCount: jsii.Number(123),
//   			McuCount: jsii.Number(123),
//   			MinWorkerCount: jsii.Number(123),
//   			ScaleInPolicy: &ScaleInPolicyProperty{
//   				CpuUtilizationPercentage: jsii.Number(123),
//   			},
//   			ScaleOutPolicy: &ScaleOutPolicyProperty{
//   				CpuUtilizationPercentage: jsii.Number(123),
//   			},
//   		},
//   		ProvisionedCapacity: &ProvisionedCapacityProperty{
//   			McuCount: jsii.Number(123),
//   			WorkerCount: jsii.Number(123),
//   		},
//   	},
//   	ConnectorConfiguration: map[string]*string{
//   		"connectorConfigurationKey": jsii.String("connectorConfiguration"),
//   	},
//   	ConnectorDescription: jsii.String("connectorDescription"),
//   	ConnectorName: jsii.String("connectorName"),
//   	KafkaCluster: &KafkaClusterProperty{
//   		ApacheKafkaCluster: &ApacheKafkaClusterProperty{
//   			BootstrapServers: jsii.String("bootstrapServers"),
//   			Vpc: &VpcProperty{
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   	},
//   	KafkaClusterClientAuthentication: &KafkaClusterClientAuthenticationProperty{
//   		AuthenticationType: jsii.String("authenticationType"),
//   	},
//   	KafkaClusterEncryptionInTransit: &KafkaClusterEncryptionInTransitProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   	},
//   	KafkaConnectVersion: jsii.String("kafkaConnectVersion"),
//   	LogDelivery: &LogDeliveryProperty{
//   		WorkerLogDelivery: &WorkerLogDeliveryProperty{
//   			CloudWatchLogs: &CloudWatchLogsLogDeliveryProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroup: jsii.String("logGroup"),
//   			},
//   			Firehose: &FirehoseLogDeliveryProperty{
//   				DeliveryStream: jsii.String("deliveryStream"),
//   				Enabled: jsii.Boolean(false),
//   			},
//   			S3: &S3LogDeliveryProperty{
//   				Bucket: jsii.String("bucket"),
//   				Enabled: jsii.Boolean(false),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	Plugins: []interface{}{
//   		&PluginProperty{
//   			CustomPlugin: &CustomPluginProperty{
//   				CustomPluginArn: jsii.String("customPluginArn"),
//   				Revision: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkerConfiguration: &WorkerConfigurationProperty{
//   		Revision: jsii.Number(123),
//   		WorkerConfigurationArn: jsii.String("workerConfigurationArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html
//
type CfnConnectorMixinProps struct {
	// The connector's compute capacity settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-capacity
	//
	Capacity interface{} `field:"optional" json:"capacity" yaml:"capacity"`
	// The configuration of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-connectorconfiguration
	//
	ConnectorConfiguration interface{} `field:"optional" json:"connectorConfiguration" yaml:"connectorConfiguration"`
	// The description of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-connectordescription
	//
	ConnectorDescription *string `field:"optional" json:"connectorDescription" yaml:"connectorDescription"`
	// The name of the connector.
	//
	// The connector name must be unique and can include up to 128 characters. Valid characters you can include in a connector name are: a-z, A-Z, 0-9, and -.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-connectorname
	//
	ConnectorName *string `field:"optional" json:"connectorName" yaml:"connectorName"`
	// The details of the Apache Kafka cluster to which the connector is connected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-kafkacluster
	//
	KafkaCluster interface{} `field:"optional" json:"kafkaCluster" yaml:"kafkaCluster"`
	// The type of client authentication used to connect to the Apache Kafka cluster.
	//
	// The value is NONE when no client authentication is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-kafkaclusterclientauthentication
	//
	KafkaClusterClientAuthentication interface{} `field:"optional" json:"kafkaClusterClientAuthentication" yaml:"kafkaClusterClientAuthentication"`
	// Details of encryption in transit to the Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-kafkaclusterencryptionintransit
	//
	KafkaClusterEncryptionInTransit interface{} `field:"optional" json:"kafkaClusterEncryptionInTransit" yaml:"kafkaClusterEncryptionInTransit"`
	// The version of Kafka Connect.
	//
	// It has to be compatible with both the Apache Kafka cluster's version and the plugins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-kafkaconnectversion
	//
	KafkaConnectVersion *string `field:"optional" json:"kafkaConnectVersion" yaml:"kafkaConnectVersion"`
	// The settings for delivering connector logs to Amazon CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-logdelivery
	//
	LogDelivery interface{} `field:"optional" json:"logDelivery" yaml:"logDelivery"`
	// Specifies which plugin to use for the connector.
	//
	// You must specify a single-element list. Amazon MSK Connect does not currently support specifying multiple plugins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-plugins
	//
	Plugins interface{} `field:"optional" json:"plugins" yaml:"plugins"`
	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon Web Services resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-serviceexecutionrolearn
	//
	ServiceExecutionRoleArn *string `field:"optional" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// A collection of tags associated with a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The worker configurations that are in use with the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-connector.html#cfn-kafkaconnect-connector-workerconfiguration
	//
	WorkerConfiguration interface{} `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
}

