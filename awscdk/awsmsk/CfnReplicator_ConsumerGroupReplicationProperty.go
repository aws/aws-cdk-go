package awsmsk


// Details about consumer group replication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   consumerGroupReplicationProperty := &ConsumerGroupReplicationProperty{
//   	ConsumerGroupsToReplicate: []*string{
//   		jsii.String("consumerGroupsToReplicate"),
//   	},
//
//   	// the properties below are optional
//   	ConsumerGroupsToExclude: []*string{
//   		jsii.String("consumerGroupsToExclude"),
//   	},
//   	DetectAndCopyNewConsumerGroups: jsii.Boolean(false),
//   	SynchroniseConsumerGroupOffsets: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-consumergroupreplication.html
//
type CfnReplicator_ConsumerGroupReplicationProperty struct {
	// List of regular expression patterns indicating the consumer groups to copy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-consumergroupreplication.html#cfn-msk-replicator-consumergroupreplication-consumergroupstoreplicate
	//
	ConsumerGroupsToReplicate *[]*string `field:"required" json:"consumerGroupsToReplicate" yaml:"consumerGroupsToReplicate"`
	// List of regular expression patterns indicating the consumer groups that should not be replicated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-consumergroupreplication.html#cfn-msk-replicator-consumergroupreplication-consumergroupstoexclude
	//
	ConsumerGroupsToExclude *[]*string `field:"optional" json:"consumerGroupsToExclude" yaml:"consumerGroupsToExclude"`
	// Enables synchronization of consumer groups to target cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-consumergroupreplication.html#cfn-msk-replicator-consumergroupreplication-detectandcopynewconsumergroups
	//
	DetectAndCopyNewConsumerGroups interface{} `field:"optional" json:"detectAndCopyNewConsumerGroups" yaml:"detectAndCopyNewConsumerGroups"`
	// Enables synchronization of consumer group offsets to target cluster.
	//
	// The translated offsets will be written to topic __consumer_offsets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-consumergroupreplication.html#cfn-msk-replicator-consumergroupreplication-synchroniseconsumergroupoffsets
	//
	SynchroniseConsumerGroupOffsets interface{} `field:"optional" json:"synchroniseConsumerGroupOffsets" yaml:"synchroniseConsumerGroupOffsets"`
}

