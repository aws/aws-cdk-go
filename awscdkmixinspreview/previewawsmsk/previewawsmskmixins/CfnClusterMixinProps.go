package previewawsmskmixins


// Properties for CfnClusterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterMixinProps := &CfnClusterMixinProps{
//   	BrokerNodeGroupInfo: &BrokerNodeGroupInfoProperty{
//   		BrokerAzDistribution: jsii.String("brokerAzDistribution"),
//   		ClientSubnets: []*string{
//   			jsii.String("clientSubnets"),
//   		},
//   		ConnectivityInfo: &ConnectivityInfoProperty{
//   			PublicAccess: &PublicAccessProperty{
//   				Type: jsii.String("type"),
//   			},
//   			VpcConnectivity: &VpcConnectivityProperty{
//   				ClientAuthentication: &VpcConnectivityClientAuthenticationProperty{
//   					Sasl: &VpcConnectivitySaslProperty{
//   						Iam: &VpcConnectivityIamProperty{
//   							Enabled: jsii.Boolean(false),
//   						},
//   						Scram: &VpcConnectivityScramProperty{
//   							Enabled: jsii.Boolean(false),
//   						},
//   					},
//   					Tls: &VpcConnectivityTlsProperty{
//   						Enabled: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   		InstanceType: jsii.String("instanceType"),
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		StorageInfo: &StorageInfoProperty{
//   			EbsStorageInfo: &EBSStorageInfoProperty{
//   				ProvisionedThroughput: &ProvisionedThroughputProperty{
//   					Enabled: jsii.Boolean(false),
//   					VolumeThroughput: jsii.Number(123),
//   				},
//   				VolumeSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ClientAuthentication: &ClientAuthenticationProperty{
//   		Sasl: &SaslProperty{
//   			Iam: &IamProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			Scram: &ScramProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   		Tls: &TlsProperty{
//   			CertificateAuthorityArnList: []*string{
//   				jsii.String("certificateAuthorityArnList"),
//   			},
//   			Enabled: jsii.Boolean(false),
//   		},
//   		Unauthenticated: &UnauthenticatedProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	ConfigurationInfo: &ConfigurationInfoProperty{
//   		Arn: jsii.String("arn"),
//   		Revision: jsii.Number(123),
//   	},
//   	CurrentVersion: jsii.String("currentVersion"),
//   	EncryptionInfo: &EncryptionInfoProperty{
//   		EncryptionAtRest: &EncryptionAtRestProperty{
//   			DataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   		},
//   		EncryptionInTransit: &EncryptionInTransitProperty{
//   			ClientBroker: jsii.String("clientBroker"),
//   			InCluster: jsii.Boolean(false),
//   		},
//   	},
//   	EnhancedMonitoring: jsii.String("enhancedMonitoring"),
//   	KafkaVersion: jsii.String("kafkaVersion"),
//   	LoggingInfo: &LoggingInfoProperty{
//   		BrokerLogs: &BrokerLogsProperty{
//   			CloudWatchLogs: &CloudWatchLogsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroup: jsii.String("logGroup"),
//   			},
//   			Firehose: &FirehoseProperty{
//   				DeliveryStream: jsii.String("deliveryStream"),
//   				Enabled: jsii.Boolean(false),
//   			},
//   			S3: &S3Property{
//   				Bucket: jsii.String("bucket"),
//   				Enabled: jsii.Boolean(false),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	NumberOfBrokerNodes: jsii.Number(123),
//   	OpenMonitoring: &OpenMonitoringProperty{
//   		Prometheus: &PrometheusProperty{
//   			JmxExporter: &JmxExporterProperty{
//   				EnabledInBroker: jsii.Boolean(false),
//   			},
//   			NodeExporter: &NodeExporterProperty{
//   				EnabledInBroker: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Rebalancing: &RebalancingProperty{
//   		Status: jsii.String("status"),
//   	},
//   	StorageMode: jsii.String("storageMode"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html
//
type CfnClusterMixinProps struct {
	// Information about the broker nodes in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-brokernodegroupinfo
	//
	BrokerNodeGroupInfo interface{} `field:"optional" json:"brokerNodeGroupInfo" yaml:"brokerNodeGroupInfo"`
	// Includes all client authentication related information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-clientauthentication
	//
	ClientAuthentication interface{} `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// The name of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Represents the configuration that you want MSK to use for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-configurationinfo
	//
	ConfigurationInfo interface{} `field:"optional" json:"configurationInfo" yaml:"configurationInfo"`
	// The version of the cluster that you want to update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-currentversion
	//
	CurrentVersion *string `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// Includes all encryption-related information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-encryptioninfo
	//
	EncryptionInfo interface{} `field:"optional" json:"encryptionInfo" yaml:"encryptionInfo"`
	// Specifies the level of monitoring for the MSK cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-enhancedmonitoring
	//
	EnhancedMonitoring *string `field:"optional" json:"enhancedMonitoring" yaml:"enhancedMonitoring"`
	// The version of Apache Kafka.
	//
	// You can use Amazon MSK to create clusters that use [supported Apache Kafka versions](https://docs.aws.amazon.com/msk/latest/developerguide/supported-kafka-versions.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-kafkaversion
	//
	KafkaVersion *string `field:"optional" json:"kafkaVersion" yaml:"kafkaVersion"`
	// Logging info details for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-logginginfo
	//
	LoggingInfo interface{} `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// The number of broker nodes in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-numberofbrokernodes
	//
	NumberOfBrokerNodes *float64 `field:"optional" json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// The settings for open monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-openmonitoring
	//
	OpenMonitoring interface{} `field:"optional" json:"openMonitoring" yaml:"openMonitoring"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-rebalancing
	//
	Rebalancing interface{} `field:"optional" json:"rebalancing" yaml:"rebalancing"`
	// This controls storage mode for supported storage tiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-storagemode
	//
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
	// An arbitrary set of tags (key-value pairs) for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

