//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateMachineImage_FromSsmParameterParameters(parameterName *string, options *SsmParameterImageOptions) error {
	if parameterName == nil {
		return fmt.Errorf("parameter parameterName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateMachineImage_GenericLinuxParameters(amiMap *map[string]*string, props *GenericLinuxImageProps) error {
	if amiMap == nil {
		return fmt.Errorf("parameter amiMap is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateMachineImage_GenericWindowsParameters(amiMap *map[string]*string, props *GenericWindowsImageProps) error {
	if amiMap == nil {
		return fmt.Errorf("parameter amiMap is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateMachineImage_LatestAmazonLinuxParameters(props *AmazonLinuxImageProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateMachineImage_LatestWindowsParameters(version WindowsVersion, props *WindowsImageProps) error {
	if version == "" {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateMachineImage_LookupParameters(props *LookupMachineImageProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

