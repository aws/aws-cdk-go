package awsmsk


// Properties for defining a `CfnTopic`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTopicProps := &CfnTopicProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	PartitionCount: jsii.Number(123),
//   	ReplicationFactor: jsii.Number(123),
//   	TopicName: jsii.String("topicName"),
//
//   	// the properties below are optional
//   	Configs: jsii.String("configs"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html
//
type CfnTopicProps struct {
	// The Amazon Resource Name (ARN) of the MSK cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html#cfn-msk-topic-clusterarn
	//
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// The number of partitions for the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html#cfn-msk-topic-partitioncount
	//
	PartitionCount *float64 `field:"required" json:"partitionCount" yaml:"partitionCount"`
	// The replication factor for the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html#cfn-msk-topic-replicationfactor
	//
	ReplicationFactor *float64 `field:"required" json:"replicationFactor" yaml:"replicationFactor"`
	// The name of the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html#cfn-msk-topic-topicname
	//
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// Base64 encoded configuration properties of the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html#cfn-msk-topic-configs
	//
	Configs *string `field:"optional" json:"configs" yaml:"configs"`
}

