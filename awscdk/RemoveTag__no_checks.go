//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RemoveTag) validateApplyTagParameters(resource ITaggable) error {
	return nil
}

func (r *jsiiProxy_RemoveTag) validateApplyTagV2Parameters(resource ITaggableV2) error {
	return nil
}

func (r *jsiiProxy_RemoveTag) validateVisitParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRemoveTagParameters(key *string, props *TagProps) error {
	return nil
}

