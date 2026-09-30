//go:build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConnectionTypePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConnectionTypePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConnectionTypePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConnectionTypePropsMixinParameters(props *CfnConnectionTypeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

