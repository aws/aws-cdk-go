//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateInitCommand_ArgvCommandParameters(argv *[]*string, options *InitCommandOptions) error {
	if argv == nil {
		return fmt.Errorf("parameter argv is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitCommand_ShellCommandParameters(shellCommand *string, options *InitCommandOptions) error {
	if shellCommand == nil {
		return fmt.Errorf("parameter shellCommand is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

