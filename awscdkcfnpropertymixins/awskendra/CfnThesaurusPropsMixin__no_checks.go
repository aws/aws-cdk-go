//go:build no_runtime_type_checking

package awskendra

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnThesaurusPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnThesaurusPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnThesaurusPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnThesaurusPropsMixinParameters(props *CfnThesaurusMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

