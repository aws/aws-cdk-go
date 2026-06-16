package awsdeadline


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   persistentVolumeConfigurationProperty := &PersistentVolumeConfigurationProperty{
//   	MountPath: jsii.String("mountPath"),
//
//   	// the properties below are optional
//   	Iops: jsii.Number(123),
//   	LastUsedTtlHours: jsii.Number(123),
//   	SizeGiB: jsii.Number(123),
//   	ThroughputMiB: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-persistentvolumeconfiguration.html
//
type CfnFleet_PersistentVolumeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-persistentvolumeconfiguration.html#cfn-deadline-fleet-persistentvolumeconfiguration-mountpath
	//
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-persistentvolumeconfiguration.html#cfn-deadline-fleet-persistentvolumeconfiguration-iops
	//
	// Default: - 3000.
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-persistentvolumeconfiguration.html#cfn-deadline-fleet-persistentvolumeconfiguration-lastusedttlhours
	//
	// Default: - 168.
	//
	LastUsedTtlHours *float64 `field:"optional" json:"lastUsedTtlHours" yaml:"lastUsedTtlHours"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-persistentvolumeconfiguration.html#cfn-deadline-fleet-persistentvolumeconfiguration-sizegib
	//
	// Default: - 250.
	//
	SizeGiB *float64 `field:"optional" json:"sizeGiB" yaml:"sizeGiB"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-persistentvolumeconfiguration.html#cfn-deadline-fleet-persistentvolumeconfiguration-throughputmib
	//
	// Default: - 125.
	//
	ThroughputMiB *float64 `field:"optional" json:"throughputMiB" yaml:"throughputMiB"`
}

