//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FeatureFlags) validateIsEnabledParameters(featureFlag *string) error {
	return nil
}

func validateFeatureFlags_OfParameters(scope constructs.IConstruct) error {
	return nil
}

