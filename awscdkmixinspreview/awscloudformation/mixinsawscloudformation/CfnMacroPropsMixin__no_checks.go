//go:build no_runtime_type_checking

package mixinsawscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMacroPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMacroPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMacroPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMacroPropsMixinParameters(props *CfnMacroMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

