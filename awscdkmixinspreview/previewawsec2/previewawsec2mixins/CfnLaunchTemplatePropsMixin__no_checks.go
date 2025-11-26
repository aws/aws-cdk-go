//go:build no_runtime_type_checking

package previewawsec2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLaunchTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLaunchTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLaunchTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLaunchTemplatePropsMixinParameters(props *CfnLaunchTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

