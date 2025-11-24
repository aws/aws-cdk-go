//go:build no_runtime_type_checking

package mixinsawsconnect

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEvaluationFormPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEvaluationFormPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEvaluationFormPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEvaluationFormPropsMixinParameters(props *CfnEvaluationFormMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

