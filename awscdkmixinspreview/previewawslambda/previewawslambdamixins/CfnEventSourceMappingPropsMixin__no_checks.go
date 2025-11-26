//go:build no_runtime_type_checking

package previewawslambdamixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEventSourceMappingPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEventSourceMappingPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEventSourceMappingPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEventSourceMappingPropsMixinParameters(props *CfnEventSourceMappingMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

