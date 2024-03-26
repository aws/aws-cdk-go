package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicReplicationProperty := &TopicReplicationProperty{
//   	TopicsToReplicate: []*string{
//   		jsii.String("topicsToReplicate"),
//   	},
//
//   	// the properties below are optional
//   	CopyAccessControlListsForTopics: jsii.Boolean(false),
//   	CopyTopicConfigurations: jsii.Boolean(false),
//   	DetectAndCopyNewTopics: jsii.Boolean(false),
//   	StartingPosition: &ReplicationStartingPositionProperty{
//   		Type: jsii.String("type"),
//   	},
//   	TopicsToExclude: []*string{
//   		jsii.String("topicsToExclude"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html
//
type CfnReplicator_TopicReplicationProperty struct {
	// List of regular expression patterns indicating the topics to copy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-topicstoreplicate
	//
	TopicsToReplicate *[]*string `field:"required" json:"topicsToReplicate" yaml:"topicsToReplicate"`
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
	// Configuration for specifying the position in the topics to start replicating from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-startingposition
	//
	StartingPosition interface{} `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// List of regular expression patterns indicating the topics that should not be replicated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-topicreplication.html#cfn-msk-replicator-topicreplication-topicstoexclude
	//
	TopicsToExclude *[]*string `field:"optional" json:"topicsToExclude" yaml:"topicsToExclude"`
}

