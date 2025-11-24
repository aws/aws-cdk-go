//go:build no_runtime_type_checking

package mixinsawsgamelift

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLocationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLocationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLocationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLocationPropsMixinParameters(props *CfnLocationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

