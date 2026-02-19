package previewawssagemakermixins


// Slurm configuration for the instance group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterSlurmConfigProperty := &ClusterSlurmConfigProperty{
//   	NodeType: jsii.String("nodeType"),
//   	PartitionNames: []*string{
//   		jsii.String("partitionNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterslurmconfig.html
//
type CfnClusterPropsMixin_ClusterSlurmConfigProperty struct {
	// The type of Slurm node for this instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterslurmconfig.html#cfn-sagemaker-cluster-clusterslurmconfig-nodetype
	//
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// The Slurm partitions that this instance group belongs to.
	//
	// Maximum of 1 partition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterslurmconfig.html#cfn-sagemaker-cluster-clusterslurmconfig-partitionnames
	//
	PartitionNames *[]*string `field:"optional" json:"partitionNames" yaml:"partitionNames"`
}

