package awsec2


// Supported EBS volume types for blockDevices.
//
// Example:
//   var vpc vpc
//
//
//   instance := ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_M5, ec2.InstanceSize_XLARGE),
//   	MachineImage: ec2.NewAmazonLinuxImage(),
//   	Vpc: vpc,
//   	HibernationEnabled: jsii.Boolean(true),
//   	BlockDevices: []blockDevice{
//   		&blockDevice{
//   			DeviceName: jsii.String("/dev/xvda"),
//   			Volume: ec2.BlockDeviceVolume_Ebs(jsii.Number(30), &EbsDeviceOptions{
//   				VolumeType: ec2.EbsDeviceVolumeType_GP3,
//   				Encrypted: jsii.Boolean(true),
//   				DeleteOnTermination: jsii.Boolean(true),
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

