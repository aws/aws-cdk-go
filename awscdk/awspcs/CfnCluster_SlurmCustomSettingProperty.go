package awspcs


// Additional settings that directly map to Slurm settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slurmCustomSettingProperty := &SlurmCustomSettingProperty{
//   	ParameterName: jsii.String("parameterName"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmcustomsetting.html
//
type CfnCluster_SlurmCustomSettingProperty struct {
	// AWS PCS supports configuration of the following Slurm parameters:.
	//
	// - For *clusters*
	//
	// - [`Prolog`](https://docs.aws.amazon.com/https://slurm.schedmd.com/slurm.conf.html#OPT_Prolog_1)
	// - [`Epilog`](https://docs.aws.amazon.com/https://slurm.schedmd.com/slurm.conf.html#OPT_Epilog_1)
	// - [`SelectTypeParameters`](https://docs.aws.amazon.com/https://slurm.schedmd.com/slurm.conf.html#OPT_SelectTypeParameters)
	// - For *compute node groups*
	//
	// - [`Weight`](https://docs.aws.amazon.com/https://slurm.schedmd.com/slurm.conf.html#OPT_Weight)
	// - [`RealMemory`](https://docs.aws.amazon.com/https://slurm.schedmd.com/slurm.conf.html#OPT_Weight)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmcustomsetting.html#cfn-pcs-cluster-slurmcustomsetting-parametername
	//
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The values for the configured Slurm settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmcustomsetting.html#cfn-pcs-cluster-slurmcustomsetting-parametervalue
	//
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

