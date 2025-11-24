//go:build no_runtime_type_checking

package mixinsawsrobomaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRobotPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRobotPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRobotPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRobotPropsMixinParameters(props *CfnRobotMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

