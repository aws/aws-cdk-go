//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func validateEcsVolume_EfsParameters(options *EfsVolumeOptions) error {
	return nil
}

func validateEcsVolume_HostParameters(options *HostVolumeOptions) error {
	return nil
}

func validateNewEcsVolumeParameters(options *EcsVolumeOptions) error {
	return nil
}

