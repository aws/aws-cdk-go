//go:build no_runtime_type_checking

package mixinsawsneptune

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDBParameterGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDBParameterGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDBParameterGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDBParameterGroupPropsMixinParameters(props *CfnDBParameterGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

