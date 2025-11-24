//go:build no_runtime_type_checking

package mixinsawsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMethodPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMethodPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMethodPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMethodPropsMixinParameters(props *CfnMethodMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

