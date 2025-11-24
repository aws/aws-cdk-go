//go:build no_runtime_type_checking

package mixinsawsomics

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkflowVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWorkflowVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWorkflowVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWorkflowVersionPropsMixinParameters(props *CfnWorkflowVersionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

