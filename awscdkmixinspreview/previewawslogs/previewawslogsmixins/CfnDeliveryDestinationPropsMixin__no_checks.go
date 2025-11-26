//go:build no_runtime_type_checking

package previewawslogsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDeliveryDestinationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDeliveryDestinationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDeliveryDestinationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDeliveryDestinationPropsMixinParameters(props *CfnDeliveryDestinationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

