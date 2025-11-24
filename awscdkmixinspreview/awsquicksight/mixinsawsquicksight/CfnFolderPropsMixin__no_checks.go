//go:build no_runtime_type_checking

package mixinsawsquicksight

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFolderPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFolderPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFolderPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFolderPropsMixinParameters(props *CfnFolderMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

