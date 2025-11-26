//go:build no_runtime_type_checking

package previewawsomicsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAnnotationStorePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAnnotationStorePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAnnotationStorePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAnnotationStorePropsMixinParameters(props *CfnAnnotationStoreMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

