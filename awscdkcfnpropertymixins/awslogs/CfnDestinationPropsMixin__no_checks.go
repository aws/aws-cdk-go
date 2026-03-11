//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDestinationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDestinationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDestinationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDestinationPropsMixinParameters(props *CfnDestinationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

