//go:build no_runtime_type_checking

package awsconfig

// Building without runtime type checking enabled, so all the below just return nil

func validateRuleScope_FromResourceParameters(resourceType ResourceType) error {
	return nil
}

func validateRuleScope_FromResourcesParameters(resourceTypes *[]ResourceType) error {
	return nil
}

func validateRuleScope_FromTagParameters(key *string) error {
	return nil
}

