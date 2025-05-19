//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func validateUserPoolOperation_OfParameters(name *string) error {
	return nil
}

