//go:build no_runtime_type_checking

package awsdevopsagent

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPrivateConnectionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPrivateConnectionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPrivateConnectionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPrivateConnectionPropsMixinParameters(props *CfnPrivateConnectionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

