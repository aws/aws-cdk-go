//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVolumeAttachmentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVolumeAttachmentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVolumeAttachmentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVolumeAttachmentPropsMixinParameters(props *CfnVolumeAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

