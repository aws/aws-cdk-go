package awsec2


// Block device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var blockDeviceVolume blockDeviceVolume
//
//   blockDevice := &blockDevice{
//   	deviceName: jsii.String("deviceName"),
//   	volume: blockDeviceVolume,
//
//   	// the properties below are optional
//   	mappingEnabled: jsii.Boolean(false),
//   }
//
type BlockDevice struct {
	// The device name exposed to the EC2 instance.
	//
	// For example, a value like `/dev/sdh`, `xvdh`.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html
	//
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Defines the block device volume, to be either an Amazon EBS volume or an ephemeral instance store volume.
	//
	// For example, a value like `BlockDeviceVolume.ebs(15)`, `BlockDeviceVolume.ephemeral(0)`.
	Volume BlockDeviceVolume `field:"required" json:"volume" yaml:"volume"`
	// If false, the device mapping will be suppressed.
	//
	// If set to false for the root device, the instance might fail the Amazon EC2 health check.
	// Amazon EC2 Auto Scaling launches a replacement instance if the instance fails the health check.
	MappingEnabled *bool `field:"optional" json:"mappingEnabled" yaml:"mappingEnabled"`
}

