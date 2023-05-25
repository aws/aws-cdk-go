//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Tags) validateAddParameters(key *string, value *string, props *TagProps) error {
	return nil
}

func (t *jsiiProxy_Tags) validateRemoveParameters(key *string, props *TagProps) error {
	return nil
}

func validateTags_OfParameters(scope constructs.IConstruct) error {
	return nil
}

