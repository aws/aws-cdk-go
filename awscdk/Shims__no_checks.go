//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateShims_AsAspectParameters(mixin constructs.IMixin) error {
	return nil
}

func validateShims_AsMixinParameters(aspect IAspect) error {
	return nil
}

