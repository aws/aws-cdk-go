package awssagemaker


// Specifies parameter(s) related to Slurm as orchestrator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterOrchestratorSlurmConfigProperty := &ClusterOrchestratorSlurmConfigProperty{
//   	SlurmConfigStrategy: jsii.String("slurmConfigStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterorchestratorslurmconfig.html
//
type CfnCluster_ClusterOrchestratorSlurmConfigProperty struct {
	// The strategy for managing Slurm configuration on the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterorchestratorslurmconfig.html#cfn-sagemaker-cluster-clusterorchestratorslurmconfig-slurmconfigstrategy
	//
	SlurmConfigStrategy *string `field:"optional" json:"slurmConfigStrategy" yaml:"slurmConfigStrategy"`
}

