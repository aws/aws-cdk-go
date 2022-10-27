//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateInitPackage_AptParameters(packageName *string, options *NamedPackageOptions) error {
	if packageName == nil {
		return fmt.Errorf("parameter packageName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitPackage_MsiParameters(location *string, options *LocationPackageOptions) error {
	if location == nil {
		return fmt.Errorf("parameter location is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitPackage_PythonParameters(packageName *string, options *NamedPackageOptions) error {
	if packageName == nil {
		return fmt.Errorf("parameter packageName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitPackage_RpmParameters(location *string, options *LocationPackageOptions) error {
	if location == nil {
		return fmt.Errorf("parameter location is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitPackage_RubyGemParameters(gemName *string, options *NamedPackageOptions) error {
	if gemName == nil {
		return fmt.Errorf("parameter gemName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitPackage_YumParameters(packageName *string, options *NamedPackageOptions) error {
	if packageName == nil {
		return fmt.Errorf("parameter packageName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewInitPackageParameters(type_ *string, versions *[]*string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if versions == nil {
		return fmt.Errorf("parameter versions is required, but nil was provided")
	}

	return nil
}

