//go:build no_runtime_type_checking

package awscdkbatchalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateEksVolume_EmptyDirParameters(options *EmptyDirVolumeOptions) error {
	return nil
}

func validateEksVolume_HostPathParameters(options *HostPathVolumeOptions) error {
	return nil
}

func validateEksVolume_SecretParameters(options *SecretPathVolumeOptions) error {
	return nil
}

func validateNewEksVolumeParameters(options *EksVolumeOptions) error {
	return nil
}

