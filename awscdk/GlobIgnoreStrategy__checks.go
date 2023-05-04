//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (g *jsiiProxy_GlobIgnoreStrategy) validateAddParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlobIgnoreStrategy) validateIgnoresParameters(absoluteFilePath *string) error {
	if absoluteFilePath == nil {
		return fmt.Errorf("parameter absoluteFilePath is required, but nil was provided")
	}

	return nil
}

func validateGlobIgnoreStrategy_DockerParameters(absoluteRootPath *string, patterns *[]*string) error {
	if absoluteRootPath == nil {
		return fmt.Errorf("parameter absoluteRootPath is required, but nil was provided")
	}

	if patterns == nil {
		return fmt.Errorf("parameter patterns is required, but nil was provided")
	}

	return nil
}

func validateGlobIgnoreStrategy_FromCopyOptionsParameters(options *CopyOptions, absoluteRootPath *string) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	if absoluteRootPath == nil {
		return fmt.Errorf("parameter absoluteRootPath is required, but nil was provided")
	}

	return nil
}

func validateGlobIgnoreStrategy_GitParameters(absoluteRootPath *string, patterns *[]*string) error {
	if absoluteRootPath == nil {
		return fmt.Errorf("parameter absoluteRootPath is required, but nil was provided")
	}

	if patterns == nil {
		return fmt.Errorf("parameter patterns is required, but nil was provided")
	}

	return nil
}

func validateGlobIgnoreStrategy_GlobParameters(absoluteRootPath *string, patterns *[]*string) error {
	if absoluteRootPath == nil {
		return fmt.Errorf("parameter absoluteRootPath is required, but nil was provided")
	}

	if patterns == nil {
		return fmt.Errorf("parameter patterns is required, but nil was provided")
	}

	return nil
}

func validateNewGlobIgnoreStrategyParameters(absoluteRootPath *string, patterns *[]*string) error {
	if absoluteRootPath == nil {
		return fmt.Errorf("parameter absoluteRootPath is required, but nil was provided")
	}

	if patterns == nil {
		return fmt.Errorf("parameter patterns is required, but nil was provided")
	}

	return nil
}

