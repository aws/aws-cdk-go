package previewawsmskmixins


// Properties for CfnTopicPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTopicMixinProps := &CfnTopicMixinProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	Configs: jsii.String("configs"),
//   	PartitionCount: jsii.Number(123),
//   	ReplicationFactor: jsii.Number(123),
//   	TopicName: jsii.String("topicName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html
//
type CfnTopicMixinProps struct {
	// The Amazon Resource Name (ARN) of the MSK cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html#cfn-msk-topic-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// Base64 encoded configuration properties of the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html#cfn-msk-topic-configs
	//
	Configs *string `field:"optional" json:"configs" yaml:"configs"`
	// The number of partitions for the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html#cfn-msk-topic-partitioncount
	//
	PartitionCount *float64 `field:"optional" json:"partitionCount" yaml:"partitionCount"`
	// The replication factor for the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html#cfn-msk-topic-replicationfactor
	//
	ReplicationFactor *float64 `field:"optional" json:"replicationFactor" yaml:"replicationFactor"`
	// The name of the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-topic.html#cfn-msk-topic-topicname
	//
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
}

