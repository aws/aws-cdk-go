//go:build no_runtime_type_checking

package awsdrs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRecoveryPlanPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRecoveryPlanPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRecoveryPlanPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRecoveryPlanPropsMixinParameters(props *CfnRecoveryPlanMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

