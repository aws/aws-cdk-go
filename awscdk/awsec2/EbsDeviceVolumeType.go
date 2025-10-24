package awsec2


// Supported EBS volume types for blockDevices.
//
// Example:
//   var vpc Vpc
//   var instanceType InstanceType
//   var machineImage IMachineImage
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	// ...
//
//   	BlockDevices: []BlockDevice{
//   		&BlockDevice{
//   			DeviceName: jsii.String("/dev/sda1"),
//   			Volume: ec2.BlockDeviceVolume_Ebs(jsii.Number(100), &EbsDeviceOptions{
//   				VolumeType: ec2.EbsDeviceVolumeType_GP3,
//   				Throughput: jsii.Number(250),
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
	// Provisioned IOPS SSD - IO2.
	EbsDeviceVolumeType_IO2 EbsDeviceVolumeType = "IO2"
	// General Purpose SSD - GP2.
	EbsDeviceVolumeType_GP2 EbsDeviceVolumeType = "GP2"
	// General Purpose SSD - GP3.
	EbsDeviceVolumeType_GP3 EbsDeviceVolumeType = "GP3"
	// Throughput Optimized HDD.
	EbsDeviceVolumeType_ST1 EbsDeviceVolumeType = "ST1"
	// Cold HDD.
	EbsDeviceVolumeType_SC1 EbsDeviceVolumeType = "SC1"
)

