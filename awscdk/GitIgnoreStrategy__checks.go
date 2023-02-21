//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (g *jsiiProxy_GitIgnoreStrategy) validateAddParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GitIgnoreStrategy) validateIgnoresParameters(absoluteFilePath *string) error {
	if absoluteFilePath == nil {
		return fmt.Errorf("parameter absoluteFilePath is required, but nil was provided")
	}

	return nil
}

func validateGitIgnoreStrategy_DockerParameters(absoluteRootPath *string, patterns *[]*string) error {
	if absoluteRootPath == nil {
		return fmt.Errorf("parameter absoluteRootPath is required, but nil was provided")
	}

	if patterns == nil {
		return fmt.Errorf("parameter patterns is required, but nil was provided")
	}

	return nil
}

func validateGitIgnoreStrategy_FromCopyOptionsParameters(options *CopyOptions, absoluteRootPath *string) error {
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

func validateGitIgnoreStrategy_GitParameters(absoluteRootPath *string, patterns *[]*string) error {
	if absoluteRootPath == nil {
		return fmt.Errorf("parameter absoluteRootPath is required, but nil was provided")
	}

	if patterns == nil {
		return fmt.Errorf("parameter patterns is required, but nil was provided")
	}

	return nil
}

func validateGitIgnoreStrategy_GlobParameters(absoluteRootPath *string, patterns *[]*string) error {
	if absoluteRootPath == nil {
		return fmt.Errorf("parameter absoluteRootPath is required, but nil was provided")
	}

	if patterns == nil {
		return fmt.Errorf("parameter patterns is required, but nil was provided")
	}

	return nil
}

func validateNewGitIgnoreStrategyParameters(absoluteRootPath *string, patterns *[]*string) error {
	if absoluteRootPath == nil {
		return fmt.Errorf("parameter absoluteRootPath is required, but nil was provided")
	}

	if patterns == nil {
		return fmt.Errorf("parameter patterns is required, but nil was provided")
	}

	return nil
}

