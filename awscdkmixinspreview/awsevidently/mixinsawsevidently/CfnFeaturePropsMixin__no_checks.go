//go:build no_runtime_type_checking

package mixinsawsevidently

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFeaturePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFeaturePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFeaturePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFeaturePropsMixinParameters(props *CfnFeatureMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

