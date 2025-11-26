package previewawsmskmixins


// Details about topic replication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   topicReplicationProperty := &TopicReplicationProperty{
//   	CopyAccessControlListsForTopics: jsii.Boolean(false),
//   	CopyTopicConfigurations: jsii.Boolean(false),
//   	DetectAndCopyNewTopics: jsii.Boolean(false),
//   	StartingPosition: &ReplicationStartingPositionProperty{
//   		Type: jsii.String("type"),
//   	},
//   	TopicNameConfiguration: &ReplicationTopicNameConfigurationProperty{
//   		Type: jsii.String("type"),
//   	},
//   	TopicsToExclude: []*string{
//   		jsii.String("topicsToExclude"),
//   	},
//   	TopicsToReplicate: []*string{
//   		jsii.String("topicsToReplicate"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html
//
type CfnReplicatorPropsMixin_TopicReplicationProperty struct {
	// Whether to periodically configure remote topic ACLs to match their corresponding upstream topics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-copyaccesscontrollistsfortopics
	//
	CopyAccessControlListsForTopics interface{} `field:"optional" json:"copyAccessControlListsForTopics" yaml:"copyAccessControlListsForTopics"`
	// Whether to periodically configure remote topics to match their corresponding upstream topics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-copytopicconfigurations
	//
	CopyTopicConfigurations interface{} `field:"optional" json:"copyTopicConfigurations" yaml:"copyTopicConfigurations"`
	// Whether to periodically check for new topics and partitions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-detectandcopynewtopics
	//
	DetectAndCopyNewTopics interface{} `field:"optional" json:"detectAndCopyNewTopics" yaml:"detectAndCopyNewTopics"`
	// Specifies the position in the topics to start replicating from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-startingposition
	//
	StartingPosition interface{} `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// Configuration for specifying replicated topic names will be the same as their corresponding upstream topics or prefixed with source cluster alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-topicnameconfiguration
	//
	TopicNameConfiguration interface{} `field:"optional" json:"topicNameConfiguration" yaml:"topicNameConfiguration"`
	// List of regular expression patterns indicating the topics that should not be replicated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-topicstoexclude
	//
	TopicsToExclude *[]*string `field:"optional" json:"topicsToExclude" yaml:"topicsToExclude"`
	// List of regular expression patterns indicating the topics to copy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-topicstoreplicate
	//
	TopicsToReplicate *[]*string `field:"optional" json:"topicsToReplicate" yaml:"topicsToReplicate"`
}

