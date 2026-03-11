//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnNetworkInterfaceAttachmentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnNetworkInterfaceAttachmentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnNetworkInterfaceAttachmentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnNetworkInterfaceAttachmentPropsMixinParameters(props *CfnNetworkInterfaceAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

