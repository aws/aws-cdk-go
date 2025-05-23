//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func validateInputFormat_CsvParameters(options *CsvOptions) error {
	return nil
}

