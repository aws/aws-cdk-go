//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEC2FleetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEC2FleetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEC2FleetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEC2FleetPropsMixinParameters(props *CfnEC2FleetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

