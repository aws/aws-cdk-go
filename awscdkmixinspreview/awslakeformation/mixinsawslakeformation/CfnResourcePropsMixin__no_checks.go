//go:build no_runtime_type_checking

package mixinsawslakeformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourcePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnResourcePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnResourcePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnResourcePropsMixinParameters(props *CfnResourceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

