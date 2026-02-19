//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func validateJsonParsingEngine_OfParameters(parsingEngine *string) error {
	return nil
}

