//go:build no_runtime_type_checking

package mixinsawsdatasync

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTaskPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTaskPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTaskPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTaskPropsMixinParameters(props *CfnTaskMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

