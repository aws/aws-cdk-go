//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnIPv4PoolPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnIPv4PoolPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnIPv4PoolPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnIPv4PoolPropsMixinParameters(props *CfnIPv4PoolMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

