package awsmsk


// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &cfnClusterProps{
//   	brokerNodeGroupInfo: &brokerNodeGroupInfoProperty{
//   		clientSubnets: []*string{
//   			jsii.String("clientSubnets"),
//   		},
//   		instanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		brokerAzDistribution: jsii.String("brokerAzDistribution"),
//   		connectivityInfo: &connectivityInfoProperty{
//   			publicAccess: &publicAccessProperty{
//   				type: jsii.String("type"),
//   			},
//   			vpcConnectivity: &vpcConnectivityProperty{
//   				clientAuthentication: &vpcConnectivityClientAuthenticationProperty{
//   					sasl: &vpcConnectivitySaslProperty{
//   						iam: &vpcConnectivityIamProperty{
//   							enabled: jsii.Boolean(false),
//   						},
//   						scram: &vpcConnectivityScramProperty{
//   							enabled: jsii.Boolean(false),
//   						},
//   					},
//   					tls: &vpcConnectivityTlsProperty{
//   						enabled: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		storageInfo: &storageInfoProperty{
//   			ebsStorageInfo: &eBSStorageInfoProperty{
//   				provisionedThroughput: &provisionedThroughputProperty{
//   					enabled: jsii.Boolean(false),
//   					volumeThroughput: jsii.Number(123),
//   				},
//   				volumeSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   	clusterName: jsii.String("clusterName"),
//   	kafkaVersion: jsii.String("kafkaVersion"),
//   	numberOfBrokerNodes: jsii.Number(123),
//
//   	// the properties below are optional
//   	clientAuthentication: &clientAuthenticationProperty{
//   		sasl: &saslProperty{
//   			iam: &iamProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   			scram: &scramProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   		tls: &tlsProperty{
//   			certificateAuthorityArnList: []*string{
//   				jsii.String("certificateAuthorityArnList"),
//   			},
//   			enabled: jsii.Boolean(false),
//   		},
//   		unauthenticated: &unauthenticatedProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	configurationInfo: &configurationInfoProperty{
//   		arn: jsii.String("arn"),
//   		revision: jsii.Number(123),
//   	},
//   	currentVersion: jsii.String("currentVersion"),
//   	encryptionInfo: &encryptionInfoProperty{
//   		encryptionAtRest: &encryptionAtRestProperty{
//   			dataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   		},
//   		encryptionInTransit: &encryptionInTransitProperty{
//   			clientBroker: jsii.String("clientBroker"),
//   			inCluster: jsii.Boolean(false),
//   		},
//   	},
//   	enhancedMonitoring: jsii.String("enhancedMonitoring"),
//   	loggingInfo: &loggingInfoProperty{
//   		brokerLogs: &brokerLogsProperty{
//   			cloudWatchLogs: &cloudWatchLogsProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				logGroup: jsii.String("logGroup"),
//   			},
//   			firehose: &firehoseProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				deliveryStream: jsii.String("deliveryStream"),
//   			},
//   			s3: &s3Property{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				bucket: jsii.String("bucket"),
//   				prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	openMonitoring: &openMonitoringProperty{
//   		prometheus: &prometheusProperty{
//   			jmxExporter: &jmxExporterProperty{
//   				enabledInBroker: jsii.Boolean(false),
//   			},
//   			nodeExporter: &nodeExporterProperty{
//   				enabledInBroker: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	storageMode: jsii.String("storageMode"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnClusterProps struct {
	// The setup to be used for brokers in the cluster.
	//
	// AWS CloudFormation may replace the cluster when you update certain `BrokerNodeGroupInfo` properties. To understand the update behavior for your use case, you should review the child properties for [`BrokerNodeGroupInfo`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#aws-properties-msk-cluster-brokernodegroupinfo-properties) .
	BrokerNodeGroupInfo interface{} `field:"required" json:"brokerNodeGroupInfo" yaml:"brokerNodeGroupInfo"`
	// The name of the cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The version of Apache Kafka.
	//
	// For more information, see [Supported Apache Kafka versions](https://docs.aws.amazon.com/msk/latest/developerguide/supported-kafka-versions.html) in the Amazon MSK Developer Guide.
	KafkaVersion *string `field:"required" json:"kafkaVersion" yaml:"kafkaVersion"`
	// The number of broker nodes you want in the Amazon MSK cluster.
	//
	// You can submit an update to increase the number of broker nodes in a cluster.
	NumberOfBrokerNodes *float64 `field:"required" json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// Includes information related to client authentication.
	ClientAuthentication interface{} `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// The Amazon MSK configuration to use for the cluster.
	ConfigurationInfo interface{} `field:"optional" json:"configurationInfo" yaml:"configurationInfo"`
	// The version of the cluster that you want to update.
	CurrentVersion *string `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// Includes all encryption-related information.
	EncryptionInfo interface{} `field:"optional" json:"encryptionInfo" yaml:"encryptionInfo"`
	// Specifies the level of monitoring for the MSK cluster.
	//
	// The possible values are `DEFAULT` , `PER_BROKER` , and `PER_TOPIC_PER_BROKER` .
	EnhancedMonitoring *string `field:"optional" json:"enhancedMonitoring" yaml:"enhancedMonitoring"`
	// You can configure your Amazon MSK cluster to send broker logs to different destination types.
	//
	// This is a container for the configuration details related to broker logs.
	LoggingInfo interface{} `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// The settings for open monitoring.
	OpenMonitoring interface{} `field:"optional" json:"openMonitoring" yaml:"openMonitoring"`
	// `AWS::MSK::Cluster.StorageMode`.
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
	// A map of key:value pairs to apply to this resource.
	//
	// Both key and value are of type String.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

