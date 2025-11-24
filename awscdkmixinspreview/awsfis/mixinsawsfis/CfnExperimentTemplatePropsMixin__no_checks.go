//go:build no_runtime_type_checking

package mixinsawsfis

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnExperimentTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnExperimentTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnExperimentTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnExperimentTemplatePropsMixinParameters(props *CfnExperimentTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

