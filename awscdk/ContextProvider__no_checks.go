//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateContextProvider_GetKeyParameters(scope constructs.Construct, options *GetContextKeyOptions) error {
	return nil
}

func validateContextProvider_GetValueParameters(scope constructs.Construct, options *GetContextValueOptions) error {
	return nil
}

