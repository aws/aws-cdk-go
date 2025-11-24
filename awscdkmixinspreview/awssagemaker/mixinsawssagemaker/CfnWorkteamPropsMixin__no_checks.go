//go:build no_runtime_type_checking

package mixinsawssagemaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkteamPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWorkteamPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWorkteamPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWorkteamPropsMixinParameters(props *CfnWorkteamMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

