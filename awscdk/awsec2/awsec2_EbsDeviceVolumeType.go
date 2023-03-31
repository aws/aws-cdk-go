package awsec2


// Supported EBS volume types for blockDevices.
//
// Example:
//   domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: es.elasticsearchVersion_V7_4(),
//   	ebs: &ebsOptions{
//   		volumeSize: jsii.Number(100),
//   		volumeType: ec2.ebsDeviceVolumeType_GENERAL_PURPOSE_SSD,
//   	},
//   	nodeToNodeEncryption: jsii.Boolean(true),
//   	encryptionAtRest: &encryptionAtRestOptions{
//   		enabled: jsii.Boolean(true),
//   	},
//   })
//
// Experimental.
type EbsDeviceVolumeType string

const (
	// Magnetic.
	// Experimental.
	EbsDeviceVolumeType_STANDARD EbsDeviceVolumeType = "STANDARD"
	// Provisioned IOPS SSD - IO1.
	// Experimental.
	EbsDeviceVolumeType_IO1 EbsDeviceVolumeType = "IO1"
	// Provisioned IOPS SSD - IO2.
	// Experimental.
	EbsDeviceVolumeType_IO2 EbsDeviceVolumeType = "IO2"
	// General Purpose SSD - GP2.
	// Experimental.
	EbsDeviceVolumeType_GP2 EbsDeviceVolumeType = "GP2"
	// General Purpose SSD - GP3.
	// Experimental.
	EbsDeviceVolumeType_GP3 EbsDeviceVolumeType = "GP3"
	// Throughput Optimized HDD.
	// Experimental.
	EbsDeviceVolumeType_ST1 EbsDeviceVolumeType = "ST1"
	// Cold HDD.
	// Experimental.
	EbsDeviceVolumeType_SC1 EbsDeviceVolumeType = "SC1"
)

