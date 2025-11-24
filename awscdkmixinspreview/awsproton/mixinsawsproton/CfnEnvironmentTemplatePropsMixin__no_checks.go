//go:build no_runtime_type_checking

package mixinsawsproton

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEnvironmentTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEnvironmentTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEnvironmentTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEnvironmentTemplatePropsMixinParameters(props *CfnEnvironmentTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

