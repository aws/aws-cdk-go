//go:build no_runtime_type_checking

package awscontrolcatalog

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnControlPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnControlPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnControlPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnControlPropsMixinParameters(props *CfnControlMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

