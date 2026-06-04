package awspcs


// Additional options related to the Slurm scheduler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slurmConfigurationProperty := &SlurmConfigurationProperty{
//   	ScaleDownIdleTimeInSeconds: jsii.Number(123),
//   	SlurmCustomSettings: []interface{}{
//   		&SlurmCustomSettingProperty{
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-slurmconfiguration.html
//
type CfnComputeNodeGroup_SlurmConfigurationProperty struct {
	// The time before an idle node is scaled down.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-slurmconfiguration.html#cfn-pcs-computenodegroup-slurmconfiguration-scaledownidletimeinseconds
	//
	ScaleDownIdleTimeInSeconds *float64 `field:"optional" json:"scaleDownIdleTimeInSeconds" yaml:"scaleDownIdleTimeInSeconds"`
	// Additional Slurm-specific configuration that directly maps to Slurm settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-slurmconfiguration.html#cfn-pcs-computenodegroup-slurmconfiguration-slurmcustomsettings
	//
	SlurmCustomSettings interface{} `field:"optional" json:"slurmCustomSettings" yaml:"slurmCustomSettings"`
}

