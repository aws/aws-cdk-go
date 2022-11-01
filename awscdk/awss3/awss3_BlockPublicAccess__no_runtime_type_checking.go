//go:build no_runtime_type_checking

package awss3

// Building without runtime type checking enabled, so all the below just return nil

func validateNewBlockPublicAccessParameters(options *BlockPublicAccessOptions) error {
	return nil
}

