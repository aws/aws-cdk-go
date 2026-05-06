//go:build no_runtime_type_checking

package awschime

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAppInstancePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAppInstancePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAppInstancePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAppInstancePropsMixinParameters(props *CfnAppInstanceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

