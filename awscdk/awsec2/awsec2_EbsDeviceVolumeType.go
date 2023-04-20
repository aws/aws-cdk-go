package awsec2


// Supported EBS volume types for blockDevices.
//
// Example:
//   domain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: es.ElasticsearchVersion_V7_4(),
//   	Ebs: &EbsOptions{
//   		VolumeSize: jsii.Number(100),
//   		VolumeType: ec2.EbsDeviceVolumeType_GENERAL_PURPOSE_SSD,
//   	},
//   	NodeToNodeEncryption: jsii.Boolean(true),
//   	EncryptionAtRest: &EncryptionAtRestOptions{
//   		Enabled: jsii.Boolean(true),
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

