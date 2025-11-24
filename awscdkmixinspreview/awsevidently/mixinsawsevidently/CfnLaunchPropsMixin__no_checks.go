//go:build no_runtime_type_checking

package mixinsawsevidently

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLaunchPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLaunchPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLaunchPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLaunchPropsMixinParameters(props *CfnLaunchMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

