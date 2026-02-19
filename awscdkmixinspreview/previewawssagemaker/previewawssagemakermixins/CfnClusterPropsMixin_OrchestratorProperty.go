package previewawssagemakermixins


// The orchestrator for a SageMaker HyperPod cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   orchestratorProperty := &OrchestratorProperty{
//   	Eks: &ClusterOrchestratorEksConfigProperty{
//   		ClusterArn: jsii.String("clusterArn"),
//   	},
//   	Slurm: &ClusterOrchestratorSlurmConfigProperty{
//   		SlurmConfigStrategy: jsii.String("slurmConfigStrategy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-orchestrator.html
//
type CfnClusterPropsMixin_OrchestratorProperty struct {
	// The configuration of the Amazon EKS orchestrator cluster for the SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-orchestrator.html#cfn-sagemaker-cluster-orchestrator-eks
	//
	Eks interface{} `field:"optional" json:"eks" yaml:"eks"`
	// Specifies parameter(s) related to Slurm as orchestrator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-orchestrator.html#cfn-sagemaker-cluster-orchestrator-slurm
	//
	Slurm interface{} `field:"optional" json:"slurm" yaml:"slurm"`
}

