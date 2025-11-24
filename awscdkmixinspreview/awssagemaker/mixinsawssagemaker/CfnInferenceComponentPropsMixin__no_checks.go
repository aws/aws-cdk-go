//go:build no_runtime_type_checking

package mixinsawssagemaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInferenceComponentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInferenceComponentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInferenceComponentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInferenceComponentPropsMixinParameters(props *CfnInferenceComponentMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

