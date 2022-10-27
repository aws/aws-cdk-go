//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateDependableTrait_GetParameters(instance IDependable) error {
	return nil
}

func validateDependableTrait_ImplementParameters(instance IDependable, trait DependableTrait) error {
	return nil
}

