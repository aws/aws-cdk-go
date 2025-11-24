//go:build no_runtime_type_checking

package mixinsawsdeadline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLimitPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLimitPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLimitPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLimitPropsMixinParameters(props *CfnLimitMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

