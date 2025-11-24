//go:build no_runtime_type_checking

package mixinsawsram

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourceSharePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnResourceSharePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnResourceSharePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnResourceSharePropsMixinParameters(props *CfnResourceShareMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

