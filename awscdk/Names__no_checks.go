//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNames_NodeUniqueIdParameters(node constructs.Node) error {
	return nil
}

func validateNames_UniqueIdParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNames_UniqueResourceNameParameters(construct constructs.IConstruct, options *UniqueResourceNameOptions) error {
	return nil
}

