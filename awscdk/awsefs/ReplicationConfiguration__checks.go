//go:build !no_runtime_type_checking

package awsefs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateReplicationConfiguration_ExistingFileSystemParameters(destinationFileSystem IFileSystem) error {
	if destinationFileSystem == nil {
		return fmt.Errorf("parameter destinationFileSystem is required, but nil was provided")
	}

	return nil
}

func validateReplicationConfiguration_OneZoneFileSystemParameters(region *string, availabilityZone *string) error {
	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	if availabilityZone == nil {
		return fmt.Errorf("parameter availabilityZone is required, but nil was provided")
	}

	return nil
}

func validateNewReplicationConfigurationParameters(options *ReplicationConfigurationProps) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

