//go:build no_runtime_type_checking

package previewawsinspectormixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAssessmentTargetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAssessmentTargetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAssessmentTargetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAssessmentTargetPropsMixinParameters(props *CfnAssessmentTargetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

