//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func validateHostVolume_EfsParameters(options *EfsVolumeOptions) error {
	return nil
}

func validateHostVolume_HostParameters(options *HostVolumeOptions) error {
	return nil
}

func validateHostVolume_IsHostVolumeParameters(x interface{}) error {
	return nil
}

func validateNewHostVolumeParameters(options *HostVolumeOptions) error {
	return nil
}

