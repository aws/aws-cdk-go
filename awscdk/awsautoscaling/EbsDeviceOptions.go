package awsautoscaling


// Block device options for an EBS volume.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//   	BlockDevices: []blockDevice{
//   		&blockDevice{
//   			DeviceName: jsii.String("gp3-volume"),
//   			Volume: autoscaling.BlockDeviceVolume_Ebs(jsii.Number(15), &EbsDeviceOptions{
//   				VolumeType: autoscaling.EbsDeviceVolumeType_GP3,
//   				Throughput: jsii.Number(125),
//   			}),
//   		},
//   	},
//   })
//
type EbsDeviceOptions struct {
	// Indicates whether to delete the volume when the instance is terminated.
	// Default: - true for Amazon EC2 Auto Scaling, false otherwise (e.g. EBS)
	//
	DeleteOnTermination *bool `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for `volumeType`: `EbsDeviceVolumeType.IO1`
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	// Default: - none, required for `EbsDeviceVolumeType.IO1`
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The throughput that the volume supports, in MiB/s Takes a minimum of 125 and maximum of 1000.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	// Default: - 125 MiB/s. Only valid on gp3 volumes.
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	// Default: `EbsDeviceVolumeType.GP2`
	//
	VolumeType EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
	// Specifies whether the EBS volume is encrypted.
	//
	// Encrypted EBS volumes can only be attached to instances that support Amazon EBS encryption.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances
	//
	// Default: false.
	//
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
}

