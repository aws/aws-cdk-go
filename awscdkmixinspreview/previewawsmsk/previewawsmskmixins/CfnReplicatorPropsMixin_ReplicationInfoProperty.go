package previewawsmskmixins


// Specifies configuration for replication between a source and target Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   replicationInfoProperty := &ReplicationInfoProperty{
//   	ConsumerGroupReplication: &ConsumerGroupReplicationProperty{
//   		ConsumerGroupsToExclude: []*string{
//   			jsii.String("consumerGroupsToExclude"),
//   		},
//   		ConsumerGroupsToReplicate: []*string{
//   			jsii.String("consumerGroupsToReplicate"),
//   		},
//   		DetectAndCopyNewConsumerGroups: jsii.Boolean(false),
//   		SynchroniseConsumerGroupOffsets: jsii.Boolean(false),
//   	},
//   	SourceKafkaClusterArn: jsii.String("sourceKafkaClusterArn"),
//   	TargetCompressionType: jsii.String("targetCompressionType"),
//   	TargetKafkaClusterArn: jsii.String("targetKafkaClusterArn"),
//   	TopicReplication: &TopicReplicationProperty{
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
//   		TopicsToReplicate: []*string{
//   			jsii.String("topicsToReplicate"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html
//
type CfnReplicatorPropsMixin_ReplicationInfoProperty struct {
	// Configuration relating to consumer group replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html#cfn-msk-replicator-replicationinfo-consumergroupreplication
	//
	ConsumerGroupReplication interface{} `field:"optional" json:"consumerGroupReplication" yaml:"consumerGroupReplication"`
	// The ARN of the source Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html#cfn-msk-replicator-replicationinfo-sourcekafkaclusterarn
	//
	SourceKafkaClusterArn *string `field:"optional" json:"sourceKafkaClusterArn" yaml:"sourceKafkaClusterArn"`
	// The compression type to use when producing records to target cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html#cfn-msk-replicator-replicationinfo-targetcompressiontype
	//
	TargetCompressionType *string `field:"optional" json:"targetCompressionType" yaml:"targetCompressionType"`
	// The ARN of the target Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html#cfn-msk-replicator-replicationinfo-targetkafkaclusterarn
	//
	TargetKafkaClusterArn *string `field:"optional" json:"targetKafkaClusterArn" yaml:"targetKafkaClusterArn"`
	// Configuration relating to topic replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationinfo.html#cfn-msk-replicator-replicationinfo-topicreplication
	//
	TopicReplication interface{} `field:"optional" json:"topicReplication" yaml:"topicReplication"`
}

