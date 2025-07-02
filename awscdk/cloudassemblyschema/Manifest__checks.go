//go:build !no_runtime_type_checking

package cloudassemblyschema

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateManifest_LoadAssemblyManifestParameters(filePath *string, options *LoadManifestOptions) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateManifest_LoadAssetManifestParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func validateManifest_LoadIntegManifestParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func validateManifest_SaveAssemblyManifestParameters(manifest *AssemblyManifest, filePath *string) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(manifest, func() string { return "parameter manifest" }); err != nil {
		return err
	}

	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func validateManifest_SaveAssetManifestParameters(manifest *AssetManifest, filePath *string) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(manifest, func() string { return "parameter manifest" }); err != nil {
		return err
	}

	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func validateManifest_SaveIntegManifestParameters(manifest *IntegManifest, filePath *string) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(manifest, func() string { return "parameter manifest" }); err != nil {
		return err
	}

	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

