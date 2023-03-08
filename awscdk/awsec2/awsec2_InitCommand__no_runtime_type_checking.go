//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateInitCommand_ArgvCommandParameters(argv *[]*string, options *InitCommandOptions) error {
	return nil
}

func validateInitCommand_ShellCommandParameters(shellCommand *string, options *InitCommandOptions) error {
	return nil
}

