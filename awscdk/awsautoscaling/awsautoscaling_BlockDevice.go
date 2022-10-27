package awsautoscaling


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
//   }
//
type BlockDevice struct {
	// The device name exposed to the EC2 instance.
	//
	// Supply a value like `/dev/sdh`, `xvdh`.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html
	//
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Defines the block device volume, to be either an Amazon EBS volume or an ephemeral instance store volume.
	//
	// Supply a value like `BlockDeviceVolume.ebs(15)`, `BlockDeviceVolume.ephemeral(0)`.
	Volume BlockDeviceVolume `field:"required" json:"volume" yaml:"volume"`
}

