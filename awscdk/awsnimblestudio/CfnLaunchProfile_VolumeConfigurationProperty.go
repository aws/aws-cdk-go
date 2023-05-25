package awsnimblestudio


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeConfigurationProperty := &VolumeConfigurationProperty{
//   	Iops: jsii.Number(123),
//   	Size: jsii.Number(123),
//   	Throughput: jsii.Number(123),
//   }
//
type CfnLaunchProfile_VolumeConfigurationProperty struct {
	// `CfnLaunchProfile.VolumeConfigurationProperty.Iops`.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// `CfnLaunchProfile.VolumeConfigurationProperty.Size`.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// `CfnLaunchProfile.VolumeConfigurationProperty.Throughput`.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
}

