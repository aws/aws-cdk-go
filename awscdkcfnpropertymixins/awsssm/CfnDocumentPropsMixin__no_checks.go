//go:build no_runtime_type_checking

package awsssm

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDocumentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDocumentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDocumentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDocumentPropsMixinParameters(props *CfnDocumentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

