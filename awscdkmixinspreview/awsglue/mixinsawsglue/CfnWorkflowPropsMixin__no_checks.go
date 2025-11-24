//go:build no_runtime_type_checking

package mixinsawsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkflowPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWorkflowPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWorkflowPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWorkflowPropsMixinParameters(props *CfnWorkflowMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

