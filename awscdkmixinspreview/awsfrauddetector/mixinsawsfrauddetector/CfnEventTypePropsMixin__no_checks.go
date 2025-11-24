//go:build no_runtime_type_checking

package mixinsawsfrauddetector

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEventTypePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEventTypePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEventTypePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEventTypePropsMixinParameters(props *CfnEventTypeMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

