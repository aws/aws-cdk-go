//go:build no_runtime_type_checking

package mixinsawsconnect

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnViewPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnViewPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnViewPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnViewPropsMixinParameters(props *CfnViewMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

