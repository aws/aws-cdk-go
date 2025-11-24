//go:build no_runtime_type_checking

package mixinsawskendra

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFaqPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFaqPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFaqPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFaqPropsMixinParameters(props *CfnFaqMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

