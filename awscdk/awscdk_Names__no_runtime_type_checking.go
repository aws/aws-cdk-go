//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNames_NodeUniqueIdParameters(node ConstructNode) error {
	return nil
}

func validateNames_UniqueIdParameters(construct constructs.Construct) error {
	return nil
}

