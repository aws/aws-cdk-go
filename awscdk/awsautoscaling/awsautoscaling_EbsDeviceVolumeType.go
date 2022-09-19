package awsautoscaling


// Supported EBS volume types for blockDevices.
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

