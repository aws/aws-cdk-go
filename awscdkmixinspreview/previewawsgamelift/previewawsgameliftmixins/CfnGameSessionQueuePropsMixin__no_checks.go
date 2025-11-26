//go:build no_runtime_type_checking

package previewawsgameliftmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGameSessionQueuePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGameSessionQueuePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGameSessionQueuePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGameSessionQueuePropsMixinParameters(props *CfnGameSessionQueueMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

