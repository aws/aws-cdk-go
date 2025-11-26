package previewawspcsmixins


// Additional settings that directly map to Slurm settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   slurmCustomSettingProperty := &SlurmCustomSettingProperty{
//   	ParameterName: jsii.String("parameterName"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-queue-slurmcustomsetting.html
//
type CfnQueuePropsMixin_SlurmCustomSettingProperty struct {
	// AWS PCS supports configuration of the Slurm parameters for queues:.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-queue-slurmcustomsetting.html#cfn-pcs-queue-slurmcustomsetting-parametername
	//
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The value for the configured Slurm setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-queue-slurmcustomsetting.html#cfn-pcs-queue-slurmcustomsetting-parametervalue
	//
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

