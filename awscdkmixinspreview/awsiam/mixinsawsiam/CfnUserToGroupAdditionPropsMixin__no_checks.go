//go:build no_runtime_type_checking

package mixinsawsiam

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUserToGroupAdditionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnUserToGroupAdditionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnUserToGroupAdditionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnUserToGroupAdditionPropsMixinParameters(props *CfnUserToGroupAdditionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

