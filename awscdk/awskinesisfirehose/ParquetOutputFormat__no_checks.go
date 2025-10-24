//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func validateNewParquetOutputFormatParameters(props *ParquetOutputFormatProps) error {
	return nil
}

