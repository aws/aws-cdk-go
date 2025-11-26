//go:build no_runtime_type_checking

package previewawstransfermixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgreementPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAgreementPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAgreementPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAgreementPropsMixinParameters(props *CfnAgreementMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

