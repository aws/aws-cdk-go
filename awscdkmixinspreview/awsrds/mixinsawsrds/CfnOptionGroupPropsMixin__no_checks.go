//go:build no_runtime_type_checking

package mixinsawsrds

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnOptionGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnOptionGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnOptionGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnOptionGroupPropsMixinParameters(props *CfnOptionGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

