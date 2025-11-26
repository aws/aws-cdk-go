//go:build no_runtime_type_checking

package previewawsomicsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSequenceStorePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSequenceStorePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSequenceStorePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSequenceStorePropsMixinParameters(props *CfnSequenceStoreMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

