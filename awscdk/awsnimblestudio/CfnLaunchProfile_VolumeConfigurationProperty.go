package awsnimblestudio


// Custom volume configuration for the root volumes that are attached to streaming sessions.
//
// This parameter is only allowed when `sessionPersistenceMode` is `ACTIVATED` .
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-volumeconfiguration.html
//
type CfnLaunchProfile_VolumeConfigurationProperty struct {
	// The number of I/O operations per second for the root volume that is attached to streaming session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-volumeconfiguration.html#cfn-nimblestudio-launchprofile-volumeconfiguration-iops
	//
	// Default: - 3000.
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The size of the root volume that is attached to the streaming session.
	//
	// The root volume size is measured in GiBs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-volumeconfiguration.html#cfn-nimblestudio-launchprofile-volumeconfiguration-size
	//
	// Default: - 500.
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The throughput to provision for the root volume that is attached to the streaming session.
	//
	// The throughput is measured in MiB/s.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-volumeconfiguration.html#cfn-nimblestudio-launchprofile-volumeconfiguration-throughput
	//
	// Default: - 125.
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
}

