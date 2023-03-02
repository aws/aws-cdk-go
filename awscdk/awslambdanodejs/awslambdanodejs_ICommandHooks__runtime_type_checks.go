//go:build !no_runtime_type_checking

package awslambdanodejs

import (
	"fmt"
)

func (i *jsiiProxy_ICommandHooks) validateAfterBundlingParameters(inputDir *string, outputDir *string) error {
	if inputDir == nil {
		return fmt.Errorf("parameter inputDir is required, but nil was provided")
	}

	if outputDir == nil {
		return fmt.Errorf("parameter outputDir is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ICommandHooks) validateBeforeBundlingParameters(inputDir *string, outputDir *string) error {
	if inputDir == nil {
		return fmt.Errorf("parameter inputDir is required, but nil was provided")
	}

	if outputDir == nil {
		return fmt.Errorf("parameter outputDir is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ICommandHooks) validateBeforeInstallParameters(inputDir *string, outputDir *string) error {
	if inputDir == nil {
		return fmt.Errorf("parameter inputDir is required, but nil was provided")
	}

	if outputDir == nil {
		return fmt.Errorf("parameter outputDir is required, but nil was provided")
	}

	return nil
}

