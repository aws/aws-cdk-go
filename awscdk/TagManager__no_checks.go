//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TagManager) validateRemoveTagParameters(key *string, priority *float64) error {
	return nil
}

func (t *jsiiProxy_TagManager) validateSetTagParameters(key *string, value *string) error {
	return nil
}

func validateTagManager_IsTaggableParameters(construct interface{}) error {
	return nil
}

func validateTagManager_IsTaggableV2Parameters(construct interface{}) error {
	return nil
}

func validateTagManager_OfParameters(construct interface{}) error {
	return nil
}

func validateNewTagManagerParameters(tagType TagType, resourceTypeName *string, options *TagManagerOptions) error {
	return nil
}

