//go:build no_runtime_type_checking

package previewawsiammixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRolePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRolePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRolePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRolePropsMixinParameters(props *CfnRoleMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

