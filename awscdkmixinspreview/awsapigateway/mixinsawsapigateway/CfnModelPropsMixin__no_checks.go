//go:build no_runtime_type_checking

package mixinsawsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnModelPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnModelPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnModelPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnModelPropsMixinParameters(props *CfnModelMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

