package awskafkaconnect


// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProps := &cfnConnectorProps{
//   	capacity: &capacityProperty{
//   		autoScaling: &autoScalingProperty{
//   			maxWorkerCount: jsii.Number(123),
//   			mcuCount: jsii.Number(123),
//   			minWorkerCount: jsii.Number(123),
//   			scaleInPolicy: &scaleInPolicyProperty{
//   				cpuUtilizationPercentage: jsii.Number(123),
//   			},
//   			scaleOutPolicy: &scaleOutPolicyProperty{
//   				cpuUtilizationPercentage: jsii.Number(123),
//   			},
//   		},
//   		provisionedCapacity: &provisionedCapacityProperty{
//   			workerCount: jsii.Number(123),
//
//   			// the properties below are optional
//   			mcuCount: jsii.Number(123),
//   		},
//   	},
//   	connectorConfiguration: map[string]*string{
//   		"connectorConfigurationKey": jsii.String("connectorConfiguration"),
//   	},
//   	connectorName: jsii.String("connectorName"),
//   	kafkaCluster: &kafkaClusterProperty{
//   		apacheKafkaCluster: &apacheKafkaClusterProperty{
//   			bootstrapServers: jsii.String("bootstrapServers"),
//   			vpc: &vpcProperty{
//   				securityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   				subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   	},
//   	kafkaClusterClientAuthentication: &kafkaClusterClientAuthenticationProperty{
//   		authenticationType: jsii.String("authenticationType"),
//   	},
//   	kafkaClusterEncryptionInTransit: &kafkaClusterEncryptionInTransitProperty{
//   		encryptionType: jsii.String("encryptionType"),
//   	},
//   	kafkaConnectVersion: jsii.String("kafkaConnectVersion"),
//   	plugins: []interface{}{
//   		&pluginProperty{
//   			customPlugin: &customPluginProperty{
//   				customPluginArn: jsii.String("customPluginArn"),
//   				revision: jsii.Number(123),
//   			},
//   		},
//   	},
//   	serviceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//
//   	// the properties below are optional
//   	connectorDescription: jsii.String("connectorDescription"),
//   	logDelivery: &logDeliveryProperty{
//   		workerLogDelivery: &workerLogDeliveryProperty{
//   			cloudWatchLogs: &cloudWatchLogsLogDeliveryProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				logGroup: jsii.String("logGroup"),
//   			},
//   			firehose: &firehoseLogDeliveryProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				deliveryStream: jsii.String("deliveryStream"),
//   			},
//   			s3: &s3LogDeliveryProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				bucket: jsii.String("bucket"),
//   				prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	workerConfiguration: &workerConfigurationProperty{
//   		revision: jsii.Number(123),
//   		workerConfigurationArn: jsii.String("workerConfigurationArn"),
//   	},
//   }
//
type CfnConnectorProps struct {
	// The connector's compute capacity settings.
	Capacity interface{} `field:"required" json:"capacity" yaml:"capacity"`
	// The configuration of the connector.
	ConnectorConfiguration interface{} `field:"required" json:"connectorConfiguration" yaml:"connectorConfiguration"`
	// The name of the connector.
	ConnectorName *string `field:"required" json:"connectorName" yaml:"connectorName"`
	// The details of the Apache Kafka cluster to which the connector is connected.
	KafkaCluster interface{} `field:"required" json:"kafkaCluster" yaml:"kafkaCluster"`
	// The type of client authentication used to connect to the Apache Kafka cluster.
	//
	// The value is NONE when no client authentication is used.
	KafkaClusterClientAuthentication interface{} `field:"required" json:"kafkaClusterClientAuthentication" yaml:"kafkaClusterClientAuthentication"`
	// Details of encryption in transit to the Apache Kafka cluster.
	KafkaClusterEncryptionInTransit interface{} `field:"required" json:"kafkaClusterEncryptionInTransit" yaml:"kafkaClusterEncryptionInTransit"`
	// The version of Kafka Connect.
	//
	// It has to be compatible with both the Apache Kafka cluster's version and the plugins.
	KafkaConnectVersion *string `field:"required" json:"kafkaConnectVersion" yaml:"kafkaConnectVersion"`
	// Specifies which plugin to use for the connector.
	//
	// You must specify a single-element list. Amazon MSK Connect does not currently support specifying multiple plugins.
	Plugins interface{} `field:"required" json:"plugins" yaml:"plugins"`
	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon Web Services resources.
	ServiceExecutionRoleArn *string `field:"required" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// The description of the connector.
	ConnectorDescription *string `field:"optional" json:"connectorDescription" yaml:"connectorDescription"`
	// The settings for delivering connector logs to Amazon CloudWatch Logs.
	LogDelivery interface{} `field:"optional" json:"logDelivery" yaml:"logDelivery"`
	// The worker configurations that are in use with the connector.
	WorkerConfiguration interface{} `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
}

