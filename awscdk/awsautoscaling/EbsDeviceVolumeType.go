package awsautoscaling


// Supported EBS volume types for blockDevices.
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
type EbsDeviceVolumeType string

const (
	// Magnetic.
	EbsDeviceVolumeType_STANDARD EbsDeviceVolumeType = "STANDARD"
	// Provisioned IOPS SSD - IO1.
	EbsDeviceVolumeType_IO1 EbsDeviceVolumeType = "IO1"
	// General Purpose SSD - GP2.
	EbsDeviceVolumeType_GP2 EbsDeviceVolumeType = "GP2"
	// General Purpose SSD - GP3.
	EbsDeviceVolumeType_GP3 EbsDeviceVolumeType = "GP3"
	// Throughput Optimized HDD.
	EbsDeviceVolumeType_ST1 EbsDeviceVolumeType = "ST1"
	// Cold HDD.
	EbsDeviceVolumeType_SC1 EbsDeviceVolumeType = "SC1"
)

