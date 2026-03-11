//go:build no_runtime_type_checking

package awsiottwinmaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEntityPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEntityPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEntityPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEntityPropsMixinParameters(props *CfnEntityMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

