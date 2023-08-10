//go:build no_runtime_type_checking

package awscdkappconfigalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateNewSnsDestinationParameters(topic awssns.ITopic) error {
	return nil
}

