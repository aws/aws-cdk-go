//go:build no_runtime_type_checking

package awskafkaconnect

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConnectorPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConnectorPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConnectorPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConnectorPropsMixinParameters(props *CfnConnectorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

