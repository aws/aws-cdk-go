package awsautoscaling


// Supported EBS volume types for blockDevices.
// Experimental.
type EbsDeviceVolumeType string

const (
	// Magnetic.
	// Experimental.
	EbsDeviceVolumeType_STANDARD EbsDeviceVolumeType = "STANDARD"
	// Provisioned IOPS SSD - IO1.
	// Experimental.
	EbsDeviceVolumeType_IO1 EbsDeviceVolumeType = "IO1"
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

