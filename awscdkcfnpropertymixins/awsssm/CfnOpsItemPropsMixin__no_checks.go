//go:build no_runtime_type_checking

package awsssm

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnOpsItemPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnOpsItemPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnOpsItemPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnOpsItemPropsMixinParameters(props *CfnOpsItemMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

