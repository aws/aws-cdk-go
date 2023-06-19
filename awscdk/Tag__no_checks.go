//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Tag) validateApplyTagParameters(resource ITaggable) error {
	return nil
}

func (t *jsiiProxy_Tag) validateVisitParameters(construct IConstruct) error {
	return nil
}

func validateTag_AddParameters(scope Construct, key *string, value *string, props *TagProps) error {
	return nil
}

func validateTag_RemoveParameters(scope Construct, key *string, props *TagProps) error {
	return nil
}

func validateNewTagParameters(key *string, value *string, props *TagProps) error {
	return nil
}

