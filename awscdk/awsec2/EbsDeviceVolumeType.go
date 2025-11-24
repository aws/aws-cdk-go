package awsec2


// Supported EBS volume types for blockDevices.
//
// Example:
//   imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("BlockDeviceImageRecipe"), &ImageRecipeProps{
//   	BaseImage: imagebuilder.BaseImage_FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
//   	BlockDevices: []BlockDevice{
//   		&BlockDevice{
//   			DeviceName: jsii.String("/dev/sda1"),
//   			Volume: ec2.BlockDeviceVolume_Ebs(jsii.Number(100), &EbsDeviceOptions{
//   				Encrypted: jsii.Boolean(true),
//   				VolumeType: ec2.EbsDeviceVolumeType_GENERAL_PURPOSE_SSD_GP3,
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

