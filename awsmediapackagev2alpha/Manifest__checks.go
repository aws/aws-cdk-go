//go:build !no_runtime_type_checking

package awsmediapackagev2alpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateManifest_DashParameters(manifest *DashManifestConfiguration) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(manifest, func() string { return "parameter manifest" }); err != nil {
		return err
	}

	return nil
}

func validateManifest_HlsParameters(manifest *HlsManifestConfiguration) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(manifest, func() string { return "parameter manifest" }); err != nil {
		return err
	}

	return nil
}

func validateManifest_LowLatencyHLSParameters(manifest *LowLatencyHlsManifestConfiguration) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(manifest, func() string { return "parameter manifest" }); err != nil {
		return err
	}

	return nil
}

func validateManifest_MssParameters(manifest *MssManifestConfiguration) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(manifest, func() string { return "parameter manifest" }); err != nil {
		return err
	}

	return nil
}

