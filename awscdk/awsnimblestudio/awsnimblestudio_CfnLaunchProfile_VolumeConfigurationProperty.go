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
//   volumeConfigurationProperty := &volumeConfigurationProperty{
//   	iops: jsii.Number(123),
//   	size: jsii.Number(123),
//   	throughput: jsii.Number(123),
//   }
//
type CfnLaunchProfile_VolumeConfigurationProperty struct {
	// The number of I/O operations per second for the root volume that is attached to streaming session.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The size of the root volume that is attached to the streaming session.
	//
	// The root volume size is measured in GiBs.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The throughput to provision for the root volume that is attached to the streaming session.
	//
	// The throughput is measured in MiB/s.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
}

