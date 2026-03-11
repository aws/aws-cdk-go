//go:build no_runtime_type_checking

package awsconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDeliveryChannelPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDeliveryChannelPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDeliveryChannelPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDeliveryChannelPropsMixinParameters(props *CfnDeliveryChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

