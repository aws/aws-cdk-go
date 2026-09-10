package awssagemaker


// Specifies parameter(s) related to Slurm as orchestrator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   clusterOrchestratorSlurmConfigProperty := &ClusterOrchestratorSlurmConfigProperty{
//   	AccountingDatabase: &ClusterAccountingDatabaseProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   		Port: jsii.Number(123),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	SlurmConfigStrategy: jsii.String("slurmConfigStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterorchestratorslurmconfig.html
//
type CfnClusterPropsMixin_ClusterOrchestratorSlurmConfigProperty struct {
	// External MySQL-compatible accounting database that a Slurm cluster's slurmdbd connects to.
	//
	// Database credentials are supplied out-of-band through the referenced Secrets Manager secret. Supported only with Continuous node provisioning.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterorchestratorslurmconfig.html#cfn-sagemaker-cluster-clusterorchestratorslurmconfig-accountingdatabase
	//
	AccountingDatabase interface{} `field:"optional" json:"accountingDatabase" yaml:"accountingDatabase"`
	// The strategy for managing Slurm configuration on the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterorchestratorslurmconfig.html#cfn-sagemaker-cluster-clusterorchestratorslurmconfig-slurmconfigstrategy
	//
	SlurmConfigStrategy *string `field:"optional" json:"slurmConfigStrategy" yaml:"slurmConfigStrategy"`
}

