package mixinsawspcs


// The Slurm configuration for the queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   slurmConfigurationProperty := &SlurmConfigurationProperty{
//   	SlurmCustomSettings: []interface{}{
//   		&SlurmCustomSettingProperty{
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-queue-slurmconfiguration.html
//
type CfnQueuePropsMixin_SlurmConfigurationProperty struct {
	// Custom Slurm parameters that directly map to Slurm configuration settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-queue-slurmconfiguration.html#cfn-pcs-queue-slurmconfiguration-slurmcustomsettings
	//
	SlurmCustomSettings interface{} `field:"optional" json:"slurmCustomSettings" yaml:"slurmCustomSettings"`
}

