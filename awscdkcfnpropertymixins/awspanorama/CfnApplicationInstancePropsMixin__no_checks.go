//go:build no_runtime_type_checking

package awspanorama

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnApplicationInstancePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationInstancePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnApplicationInstancePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnApplicationInstancePropsMixinParameters(props *CfnApplicationInstanceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

