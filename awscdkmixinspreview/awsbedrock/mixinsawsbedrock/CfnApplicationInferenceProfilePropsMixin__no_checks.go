//go:build no_runtime_type_checking

package mixinsawsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnApplicationInferenceProfilePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationInferenceProfilePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnApplicationInferenceProfilePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnApplicationInferenceProfilePropsMixinParameters(props *CfnApplicationInferenceProfileMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

