//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPCGatewayAttachmentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVPCGatewayAttachmentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVPCGatewayAttachmentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVPCGatewayAttachmentPropsMixinParameters(props *CfnVPCGatewayAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

