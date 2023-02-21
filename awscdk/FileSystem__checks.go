//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateFileSystem_CopyDirectoryParameters(srcDir *string, destDir *string, options *CopyOptions) error {
	if srcDir == nil {
		return fmt.Errorf("parameter srcDir is required, but nil was provided")
	}

	if destDir == nil {
		return fmt.Errorf("parameter destDir is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateFileSystem_FingerprintParameters(fileOrDirectory *string, options *FingerprintOptions) error {
	if fileOrDirectory == nil {
		return fmt.Errorf("parameter fileOrDirectory is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateFileSystem_IsEmptyParameters(dir *string) error {
	if dir == nil {
		return fmt.Errorf("parameter dir is required, but nil was provided")
	}

	return nil
}

func validateFileSystem_MkdtempParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

