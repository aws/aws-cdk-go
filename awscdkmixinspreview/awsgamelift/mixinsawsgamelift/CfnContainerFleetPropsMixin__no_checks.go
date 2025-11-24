//go:build no_runtime_type_checking

package mixinsawsgamelift

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnContainerFleetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnContainerFleetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnContainerFleetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnContainerFleetPropsMixinParameters(props *CfnContainerFleetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

