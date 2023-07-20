package awsmsk


// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &CfnClusterProps{
//   	BrokerNodeGroupInfo: &BrokerNodeGroupInfoProperty{
//   		ClientSubnets: []*string{
//   			jsii.String("clientSubnets"),
//   		},
//   		InstanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		BrokerAzDistribution: jsii.String("brokerAzDistribution"),
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
//   	ClusterName: jsii.String("clusterName"),
//   	KafkaVersion: jsii.String("kafkaVersion"),
//   	NumberOfBrokerNodes: jsii.Number(123),
//
//   	// the properties below are optional
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
//   	LoggingInfo: &LoggingInfoProperty{
//   		BrokerLogs: &BrokerLogsProperty{
//   			CloudWatchLogs: &CloudWatchLogsProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				LogGroup: jsii.String("logGroup"),
//   			},
//   			Firehose: &FirehoseProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				DeliveryStream: jsii.String("deliveryStream"),
//   			},
//   			S3: &S3Property{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				Bucket: jsii.String("bucket"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
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
//   	StorageMode: jsii.String("storageMode"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html
//
type CfnClusterProps struct {
	// Information about the broker nodes in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-brokernodegroupinfo
	//
	BrokerNodeGroupInfo interface{} `field:"required" json:"brokerNodeGroupInfo" yaml:"brokerNodeGroupInfo"`
	// The name of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The version of Apache Kafka.
	//
	// You can use Amazon MSK to create clusters that use Apache Kafka versions 1.1.1 and 2.2.1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-kafkaversion
	//
	KafkaVersion *string `field:"required" json:"kafkaVersion" yaml:"kafkaVersion"`
	// The number of broker nodes in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-numberofbrokernodes
	//
	NumberOfBrokerNodes *float64 `field:"required" json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// Includes all client authentication related information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-clientauthentication
	//
	ClientAuthentication interface{} `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
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
	//
	// The possible values are `DEFAULT` , `PER_BROKER` , and `PER_TOPIC_PER_BROKER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-enhancedmonitoring
	//
	EnhancedMonitoring *string `field:"optional" json:"enhancedMonitoring" yaml:"enhancedMonitoring"`
	// Logging Info details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-logginginfo
	//
	LoggingInfo interface{} `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// The settings for open monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-openmonitoring
	//
	OpenMonitoring interface{} `field:"optional" json:"openMonitoring" yaml:"openMonitoring"`
	// This controls storage mode for supported storage tiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-storagemode
	//
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
	// Create tags when creating the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

