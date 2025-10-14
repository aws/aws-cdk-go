package awspcs


// Additional settings that directly map to Slurm settings.
//
// > AWS PCS supports a subset of Slurm settings. For more information, see [Configuring custom Slurm settings in AWS PCS](https://docs.aws.amazon.com//pcs/latest/userguide/slurm-custom-settings.html) in the *AWS PCS User Guide* .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-slurmcustomsetting.html
//
type CfnComputeNodeGroup_SlurmCustomSettingProperty struct {
	// AWS PCS supports custom Slurm settings for clusters, compute node groups, and queues.
	//
	// For more information, see [Configuring custom Slurm settings in AWS PCS](https://docs.aws.amazon.com//pcs/latest/userguide/slurm-custom-settings.html) in the *AWS PCS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-slurmcustomsetting.html#cfn-pcs-computenodegroup-slurmcustomsetting-parametername
	//
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The values for the configured Slurm settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-slurmcustomsetting.html#cfn-pcs-computenodegroup-slurmcustomsetting-parametervalue
	//
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

