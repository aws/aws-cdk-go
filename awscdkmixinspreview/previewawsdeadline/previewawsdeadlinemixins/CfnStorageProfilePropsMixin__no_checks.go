//go:build no_runtime_type_checking

package previewawsdeadlinemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStorageProfilePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStorageProfilePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStorageProfilePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStorageProfilePropsMixinParameters(props *CfnStorageProfileMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

