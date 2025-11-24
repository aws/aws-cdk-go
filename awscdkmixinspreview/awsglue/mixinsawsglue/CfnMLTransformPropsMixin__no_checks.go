//go:build no_runtime_type_checking

package mixinsawsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMLTransformPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMLTransformPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMLTransformPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMLTransformPropsMixinParameters(props *CfnMLTransformMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

