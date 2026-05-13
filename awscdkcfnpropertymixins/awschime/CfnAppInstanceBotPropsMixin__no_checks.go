//go:build no_runtime_type_checking

package awschime

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAppInstanceBotPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAppInstanceBotPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAppInstanceBotPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAppInstanceBotPropsMixinParameters(props *CfnAppInstanceBotMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

