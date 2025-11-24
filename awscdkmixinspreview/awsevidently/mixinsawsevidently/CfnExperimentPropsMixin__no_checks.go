//go:build no_runtime_type_checking

package mixinsawsevidently

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnExperimentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnExperimentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnExperimentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnExperimentPropsMixinParameters(props *CfnExperimentMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

