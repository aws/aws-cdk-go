//go:build no_runtime_type_checking

package regioninfo

// Building without runtime type checking enabled, so all the below just return nil

func validateFact_FindParameters(region *string, name *string) error {
	return nil
}

func validateFact_RegisterParameters(fact IFact) error {
	return nil
}

func validateFact_RequireFactParameters(region *string, name *string) error {
	return nil
}

func validateFact_UnregisterParameters(region *string, name *string) error {
	return nil
}

