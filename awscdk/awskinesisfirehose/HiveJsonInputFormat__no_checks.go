//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func validateNewHiveJsonInputFormatParameters(props *HiveJsonInputFormatProps) error {
	return nil
}

