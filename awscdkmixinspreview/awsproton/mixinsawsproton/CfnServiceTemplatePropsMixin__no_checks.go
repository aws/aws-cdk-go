//go:build no_runtime_type_checking

package mixinsawsproton

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnServiceTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnServiceTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnServiceTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnServiceTemplatePropsMixinParameters(props *CfnServiceTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

