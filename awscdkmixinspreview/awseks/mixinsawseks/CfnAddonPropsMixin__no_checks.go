//go:build no_runtime_type_checking

package mixinsawseks

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAddonPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAddonPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAddonPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAddonPropsMixinParameters(props *CfnAddonMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

