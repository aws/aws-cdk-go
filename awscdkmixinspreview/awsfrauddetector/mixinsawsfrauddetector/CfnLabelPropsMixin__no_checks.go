//go:build no_runtime_type_checking

package mixinsawsfrauddetector

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLabelPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLabelPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLabelPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLabelPropsMixinParameters(props *CfnLabelMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

