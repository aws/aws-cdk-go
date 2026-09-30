//go:build no_runtime_type_checking

package awschime

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVoiceConnectorPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVoiceConnectorPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVoiceConnectorPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVoiceConnectorPropsMixinParameters(props *CfnVoiceConnectorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

