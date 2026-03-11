//go:build no_runtime_type_checking

package awsiottwinmaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnComponentTypePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnComponentTypePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnComponentTypePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnComponentTypePropsMixinParameters(props *CfnComponentTypeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

