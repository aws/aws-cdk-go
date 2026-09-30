package awspcs


// Additional options related to the Slurm scheduler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slurmConfigurationProperty := &SlurmConfigurationProperty{
//   	GresCustomSettings: []interface{}{
//   		map[string]*string{
//   			"gresCustomSettingsKey": jsii.String("gresCustomSettings"),
//   		},
//   	},
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
	// Additional Slurm gres.conf records for the compute node group. Each item is a map of gres.conf attribute names to values describing one gres.conf record (for example a GPU topology, MIG, MPS, or custom GRES entry). AWS PCS adds the NodeName= prefix and merges these records with the GPU record it derives from the instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-slurmconfiguration.html#cfn-pcs-computenodegroup-slurmconfiguration-grescustomsettings
	//
	GresCustomSettings interface{} `field:"optional" json:"gresCustomSettings" yaml:"gresCustomSettings"`
	// The time before an idle node is scaled down.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-slurmconfiguration.html#cfn-pcs-computenodegroup-slurmconfiguration-scaledownidletimeinseconds
	//
	ScaleDownIdleTimeInSeconds *float64 `field:"optional" json:"scaleDownIdleTimeInSeconds" yaml:"scaleDownIdleTimeInSeconds"`
	// Additional Slurm-specific configuration that directly maps to Slurm settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-slurmconfiguration.html#cfn-pcs-computenodegroup-slurmconfiguration-slurmcustomsettings
	//
	SlurmCustomSettings interface{} `field:"optional" json:"slurmCustomSettings" yaml:"slurmCustomSettings"`
}

