//go:build no_runtime_type_checking

package awsfis

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSafetyLeverPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSafetyLeverPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSafetyLeverPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSafetyLeverPropsMixinParameters(props *CfnSafetyLeverMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

