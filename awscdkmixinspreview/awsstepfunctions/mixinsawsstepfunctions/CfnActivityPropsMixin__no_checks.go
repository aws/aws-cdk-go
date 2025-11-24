//go:build no_runtime_type_checking

package mixinsawsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnActivityPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnActivityPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnActivityPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnActivityPropsMixinParameters(props *CfnActivityMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

