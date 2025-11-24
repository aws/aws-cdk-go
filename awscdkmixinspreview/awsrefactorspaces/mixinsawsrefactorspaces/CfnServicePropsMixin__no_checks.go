//go:build no_runtime_type_checking

package mixinsawsrefactorspaces

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnServicePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnServicePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnServicePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnServicePropsMixinParameters(props *CfnServiceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

