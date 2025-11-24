//go:build no_runtime_type_checking

package mixinsawsiot

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMitigationActionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMitigationActionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMitigationActionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMitigationActionPropsMixinParameters(props *CfnMitigationActionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

