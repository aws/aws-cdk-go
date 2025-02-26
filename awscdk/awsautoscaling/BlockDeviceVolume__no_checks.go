//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func validateBlockDeviceVolume_EbsParameters(volumeSize *float64, options *EbsDeviceOptions) error {
	return nil
}

func validateBlockDeviceVolume_EbsFromSnapshotParameters(snapshotId *string, options *EbsDeviceSnapshotOptions) error {
	return nil
}

func validateBlockDeviceVolume_EphemeralParameters(volumeIndex *float64) error {
	return nil
}

func validateNewBlockDeviceVolumeParameters(ebsDevice *EbsDeviceProps) error {
	return nil
}

