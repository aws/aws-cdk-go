//go:build no_runtime_type_checking

package awsssmcontacts

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnContactPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnContactPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnContactPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnContactPropsMixinParameters(props *CfnContactMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

