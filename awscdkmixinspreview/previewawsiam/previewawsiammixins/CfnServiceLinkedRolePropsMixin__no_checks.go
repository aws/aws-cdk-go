//go:build no_runtime_type_checking

package previewawsiammixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnServiceLinkedRolePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnServiceLinkedRolePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnServiceLinkedRolePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnServiceLinkedRolePropsMixinParameters(props *CfnServiceLinkedRoleMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

