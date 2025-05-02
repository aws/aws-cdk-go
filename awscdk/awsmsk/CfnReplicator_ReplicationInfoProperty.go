package awsmsk


// Specifies configuration for replication between a source and target Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationInfoProperty := &ReplicationInfoProperty{
//   	ConsumerGroupReplication: &ConsumerGroupReplicationProperty{
//   		ConsumerGroupsToReplicate: []*string{
//   			jsii.String("consumerGroupsToReplicate"),
//   		},
//
//   		// the properties below are optional
//   		ConsumerGroupsToExclude: []*string{
//   			jsii.String("consumerGroupsToExclude"),
//   		},
//   		DetectAndCopyNewConsumerGroups: jsii.Boolean(false),
//   		SynchroniseConsumerGroupOffsets: jsii.Boolean(false),
//   	},
//   	SourceKafkaClusterArn: jsii.String("sourceKafkaClusterArn"),
//   	TargetCompressionType: jsii.String("targetCompressionType"),
//   	TargetKafkaClusterArn: jsii.String("targetKafkaClusterArn"),
//   	TopicReplication: &TopicReplicationProperty{
//   		TopicsToReplicate: []*string{
//   			jsii.String("topicsToReplicate"),
//   		},
//
//   		// the properties below are optional
//   		CopyAccessControlListsForTopics: jsii.Boolean(false),
//   		CopyTopicConfigurations: jsii.Boolean(false),
//   		DetectAndCopyNewTopics: jsii.Boolean(false),
//   		StartingPosition: &ReplicationStartingPositionProperty{
//   			Type: jsii.String("type"),
//   		},
//   		TopicNameConfiguration: &ReplicationTopicNameConfigurationProperty{
//   			Type: jsii.String("type"),
//   		},
//   		TopicsToExclude: []*string{
//   			jsii.String("topicsToExclude"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html
//
type CfnReplicator_ReplicationInfoProperty struct {
	// Configuration relating to consumer group replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html#cfn-msk-replicator-replicationinfo-consumergroupreplication
	//
	ConsumerGroupReplication interface{} `field:"required" json:"consumerGroupReplication" yaml:"consumerGroupReplication"`
	// The ARN of the source Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html#cfn-msk-replicator-replicationinfo-sourcekafkaclusterarn
	//
	SourceKafkaClusterArn *string `field:"required" json:"sourceKafkaClusterArn" yaml:"sourceKafkaClusterArn"`
	// The compression type to use when producing records to target cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html#cfn-msk-replicator-replicationinfo-targetcompressiontype
	//
	TargetCompressionType *string `field:"required" json:"targetCompressionType" yaml:"targetCompressionType"`
	// The ARN of the target Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html#cfn-msk-replicator-replicationinfo-targetkafkaclusterarn
	//
	TargetKafkaClusterArn *string `field:"required" json:"targetKafkaClusterArn" yaml:"targetKafkaClusterArn"`
	// Configuration relating to topic replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html#cfn-msk-replicator-replicationinfo-topicreplication
	//
	TopicReplication interface{} `field:"required" json:"topicReplication" yaml:"topicReplication"`
}

