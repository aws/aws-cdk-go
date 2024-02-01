package awsstepfunctionstasks


// EBS Volume Types.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_VolumeSpecification.html#EMR-Type-VolumeSpecification-VolumeType
//
type EmrCreateCluster_EbsBlockDeviceVolumeType string

const (
	// gp3 Volume Type.
	EmrCreateCluster_EbsBlockDeviceVolumeType_GP3 EmrCreateCluster_EbsBlockDeviceVolumeType = "GP3"
	// gp2 Volume Type.
	EmrCreateCluster_EbsBlockDeviceVolumeType_GP2 EmrCreateCluster_EbsBlockDeviceVolumeType = "GP2"
	// io1 Volume Type.
	EmrCreateCluster_EbsBlockDeviceVolumeType_IO1 EmrCreateCluster_EbsBlockDeviceVolumeType = "IO1"
	// st1 Volume Type.
	EmrCreateCluster_EbsBlockDeviceVolumeType_ST1 EmrCreateCluster_EbsBlockDeviceVolumeType = "ST1"
	// sc1 Volume Type.
	EmrCreateCluster_EbsBlockDeviceVolumeType_SC1 EmrCreateCluster_EbsBlockDeviceVolumeType = "SC1"
	// Standard Volume Type.
	EmrCreateCluster_EbsBlockDeviceVolumeType_STANDARD EmrCreateCluster_EbsBlockDeviceVolumeType = "STANDARD"
)

