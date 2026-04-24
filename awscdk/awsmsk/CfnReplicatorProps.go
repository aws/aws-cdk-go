package awsmsk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnReplicator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicatorProps := &CfnReplicatorProps{
//   	KafkaClusters: []interface{}{
//   		&KafkaClusterProperty{
//   			AmazonMskCluster: &AmazonMskClusterProperty{
//   				MskClusterArn: jsii.String("mskClusterArn"),
//   			},
//   			ApacheKafkaCluster: &ApacheKafkaClusterProperty{
//   				ApacheKafkaClusterId: jsii.String("apacheKafkaClusterId"),
//   				BootstrapBrokerString: jsii.String("bootstrapBrokerString"),
//   			},
//   			ClientAuthentication: &KafkaClusterClientAuthenticationProperty{
//   				SaslScram: &KafkaClusterSaslScramAuthenticationProperty{
//   					Mechanism: jsii.String("mechanism"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   			EncryptionInTransit: &KafkaClusterEncryptionInTransitProperty{
//   				EncryptionType: jsii.String("encryptionType"),
//
//   				// the properties below are optional
//   				RootCaCertificate: jsii.String("rootCaCertificate"),
//   			},
//   			VpcConfig: &KafkaClusterClientVpcConfigProperty{
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//
//   				// the properties below are optional
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   			},
//   		},
//   	},
//   	ReplicationInfoList: []interface{}{
//   		&ReplicationInfoProperty{
//   			ConsumerGroupReplication: &ConsumerGroupReplicationProperty{
//   				ConsumerGroupsToReplicate: []*string{
//   					jsii.String("consumerGroupsToReplicate"),
//   				},
//
//   				// the properties below are optional
//   				ConsumerGroupOffsetSyncMode: jsii.String("consumerGroupOffsetSyncMode"),
//   				ConsumerGroupsToExclude: []*string{
//   					jsii.String("consumerGroupsToExclude"),
//   				},
//   				DetectAndCopyNewConsumerGroups: jsii.Boolean(false),
//   				SynchroniseConsumerGroupOffsets: jsii.Boolean(false),
//   			},
//   			TargetCompressionType: jsii.String("targetCompressionType"),
//   			TopicReplication: &TopicReplicationProperty{
//   				TopicsToReplicate: []*string{
//   					jsii.String("topicsToReplicate"),
//   				},
//
//   				// the properties below are optional
//   				CopyAccessControlListsForTopics: jsii.Boolean(false),
//   				CopyTopicConfigurations: jsii.Boolean(false),
//   				DetectAndCopyNewTopics: jsii.Boolean(false),
//   				StartingPosition: &ReplicationStartingPositionProperty{
//   					Type: jsii.String("type"),
//   				},
//   				TopicNameConfiguration: &ReplicationTopicNameConfigurationProperty{
//   					Type: jsii.String("type"),
//   				},
//   				TopicsToExclude: []*string{
//   					jsii.String("topicsToExclude"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			SourceKafkaClusterArn: jsii.String("sourceKafkaClusterArn"),
//   			SourceKafkaClusterId: jsii.String("sourceKafkaClusterId"),
//   			TargetKafkaClusterArn: jsii.String("targetKafkaClusterArn"),
//   			TargetKafkaClusterId: jsii.String("targetKafkaClusterId"),
//   		},
//   	},
//   	ReplicatorName: jsii.String("replicatorName"),
//   	ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	LogDelivery: &LogDeliveryProperty{
//   		ReplicatorLogDelivery: &ReplicatorLogDeliveryProperty{
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-replicator.html
//
type CfnReplicatorProps struct {
	// Kafka Clusters to use in setting up sources / targets for replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-replicator.html#cfn-msk-replicator-kafkaclusters
	//
	KafkaClusters interface{} `field:"required" json:"kafkaClusters" yaml:"kafkaClusters"`
	// A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-replicator.html#cfn-msk-replicator-replicationinfolist
	//
	ReplicationInfoList interface{} `field:"required" json:"replicationInfoList" yaml:"replicationInfoList"`
	// The name of the replicator.
	//
	// Alpha-numeric characters with '-' are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-replicator.html#cfn-msk-replicator-replicatorname
	//
	ReplicatorName *string `field:"required" json:"replicatorName" yaml:"replicatorName"`
	// The ARN of the IAM role used by the replicator to access resources in the customer's account (e.g source and target clusters).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-replicator.html#cfn-msk-replicator-serviceexecutionrolearn
	//
	ServiceExecutionRoleArn *string `field:"required" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// A summary description of the replicator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-replicator.html#cfn-msk-replicator-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration for log delivery for the replicator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-replicator.html#cfn-msk-replicator-logdelivery
	//
	LogDelivery interface{} `field:"optional" json:"logDelivery" yaml:"logDelivery"`
	// List of tags to attach to created Replicator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-replicator.html#cfn-msk-replicator-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

