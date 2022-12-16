package awsautoscaling


// Block device options for an EBS volume.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: instanceType,
//   	machineImage: machineImage,
//   	blockDevices: []blockDevice{
//   		&blockDevice{
//   			deviceName: jsii.String("gp3-volume"),
//   			volume: autoscaling.blockDeviceVolume.ebs(jsii.Number(15), &ebsDeviceOptions{
//   				volumeType: autoscaling.ebsDeviceVolumeType_GP3,
//   				throughput: jsii.Number(125),
//   			}),
//   		},
//   	},
//   })
//
type EbsDeviceOptions struct {
	// Indicates whether to delete the volume when the instance is terminated.
	DeleteOnTermination *bool `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for {@link volumeType}: {@link EbsDeviceVolumeType.IO1}
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The throughput that the volume supports, in MiB/s Takes a minimum of 125 and maximum of 1000.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	VolumeType EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
	// Specifies whether the EBS volume is encrypted.
	//
	// Encrypted EBS volumes can only be attached to instances that support Amazon EBS encryption.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances
	//
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
}

