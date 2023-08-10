package awsnimblestudio


// <p>Custom volume configuration for the root volumes that are attached to streaming             sessions.</p>          <p>This parameter is only allowed when <code>sessionPersistenceMode</code> is                 <code>ACTIVATED</code>.</p>.
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
	// <p>The number of I/O operations per second for the root volume that is attached to             streaming session.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-volumeconfiguration.html#cfn-nimblestudio-launchprofile-volumeconfiguration-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// <p>The size of the root volume that is attached to the streaming session.
	//
	// The root volume
	//             size is measured in GiBs.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-volumeconfiguration.html#cfn-nimblestudio-launchprofile-volumeconfiguration-size
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// <p>The throughput to provision for the root volume that is attached to the streaming             session.
	//
	// The throughput is measured in MiB/s.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-volumeconfiguration.html#cfn-nimblestudio-launchprofile-volumeconfiguration-throughput
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
}

