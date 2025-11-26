//go:build no_runtime_type_checking

package previewawsomicsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnReferenceStorePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnReferenceStorePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnReferenceStorePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnReferenceStorePropsMixinParameters(props *CfnReferenceStoreMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

