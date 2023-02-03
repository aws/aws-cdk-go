package awsautoscaling


// Supported EBS volume types for blockDevices.
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

