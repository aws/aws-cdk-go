//go:build no_runtime_type_checking

package mixinsawsnetworkmanager

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVpcAttachmentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVpcAttachmentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVpcAttachmentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVpcAttachmentPropsMixinParameters(props *CfnVpcAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

