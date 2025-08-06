package awsmsk


// Information about Kafka Cluster to be used as source / target for replication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterProperty := &KafkaClusterProperty{
//   	AmazonMskCluster: &AmazonMskClusterProperty{
//   		MskClusterArn: jsii.String("mskClusterArn"),
//   	},
//   	VpcConfig: &KafkaClusterClientVpcConfigProperty{
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//
//   		// the properties below are optional
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkacluster.html
//
type CfnReplicator_KafkaClusterProperty struct {
	// Details of an Amazon MSK Cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkacluster.html#cfn-msk-replicator-kafkacluster-amazonmskcluster
	//
	AmazonMskCluster interface{} `field:"required" json:"amazonMskCluster" yaml:"amazonMskCluster"`
	// Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkacluster.html#cfn-msk-replicator-kafkacluster-vpcconfig
	//
	VpcConfig interface{} `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
}

