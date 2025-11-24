//go:build no_runtime_type_checking

package mixinsawsdeadline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFarmPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFarmPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFarmPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFarmPropsMixinParameters(props *CfnFarmMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

