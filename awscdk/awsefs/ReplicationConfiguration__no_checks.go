//go:build no_runtime_type_checking

package awsefs

// Building without runtime type checking enabled, so all the below just return nil

func validateReplicationConfiguration_ExistingFileSystemParameters(destinationFileSystem IFileSystem) error {
	return nil
}

func validateReplicationConfiguration_OneZoneFileSystemParameters(region *string, availabilityZone *string) error {
	return nil
}

func validateNewReplicationConfigurationParameters(options *ReplicationConfigurationProps) error {
	return nil
}

