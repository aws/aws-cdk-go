//go:build no_runtime_type_checking

package previewawsioteventsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDetectorModelPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDetectorModelPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDetectorModelPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDetectorModelPropsMixinParameters(props *CfnDetectorModelMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

