//go:build !no_runtime_type_checking

package awsautoscaling

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateBlockDeviceVolume_EbsParameters(volumeSize *float64, options *EbsDeviceOptions) error {
	if volumeSize == nil {
		return fmt.Errorf("parameter volumeSize is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateBlockDeviceVolume_EbsFromSnapshotParameters(snapshotId *string, options *EbsDeviceSnapshotOptions) error {
	if snapshotId == nil {
		return fmt.Errorf("parameter snapshotId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateBlockDeviceVolume_EphemeralParameters(volumeIndex *float64) error {
	if volumeIndex == nil {
		return fmt.Errorf("parameter volumeIndex is required, but nil was provided")
	}

	return nil
}

func validateNewBlockDeviceVolumeParameters(ebsDevice *EbsDeviceProps) error {
	if err := _jsii_.ValidateStruct(ebsDevice, func() string { return "parameter ebsDevice" }); err != nil {
		return err
	}

	return nil
}

