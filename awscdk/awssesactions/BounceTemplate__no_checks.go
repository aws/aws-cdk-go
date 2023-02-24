//go:build no_runtime_type_checking

package awssesactions

// Building without runtime type checking enabled, so all the below just return nil

func validateNewBounceTemplateParameters(props *BounceTemplateProps) error {
	return nil
}

