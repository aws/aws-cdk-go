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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-queue-slurmcustomsetting.html
//
type CfnQueue_SlurmCustomSettingProperty struct {
	// AWS PCS supports configuration of the Slurm parameters for queues:.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-queue-slurmcustomsetting.html#cfn-pcs-queue-slurmcustomsetting-parametername
	//
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The value for the configured Slurm setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-queue-slurmcustomsetting.html#cfn-pcs-queue-slurmcustomsetting-parametervalue
	//
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

