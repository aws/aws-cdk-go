//go:build no_runtime_type_checking

package previewawsinspectormixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAssessmentTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAssessmentTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAssessmentTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAssessmentTemplatePropsMixinParameters(props *CfnAssessmentTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

