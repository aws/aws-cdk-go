//go:build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSchemaPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSchemaPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSchemaPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSchemaPropsMixinParameters(props *CfnSchemaMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

