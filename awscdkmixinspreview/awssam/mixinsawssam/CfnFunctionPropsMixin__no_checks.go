//go:build no_runtime_type_checking

package mixinsawssam

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFunctionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFunctionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFunctionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFunctionPropsMixinParameters(props *CfnFunctionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

