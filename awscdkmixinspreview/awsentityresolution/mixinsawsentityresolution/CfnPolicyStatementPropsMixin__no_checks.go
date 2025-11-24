//go:build no_runtime_type_checking

package mixinsawsentityresolution

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPolicyStatementPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPolicyStatementPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPolicyStatementPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPolicyStatementPropsMixinParameters(props *CfnPolicyStatementMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

