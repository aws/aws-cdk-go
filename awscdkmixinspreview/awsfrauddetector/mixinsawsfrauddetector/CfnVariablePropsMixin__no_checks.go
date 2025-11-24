//go:build no_runtime_type_checking

package mixinsawsfrauddetector

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVariablePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVariablePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVariablePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVariablePropsMixinParameters(props *CfnVariableMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

