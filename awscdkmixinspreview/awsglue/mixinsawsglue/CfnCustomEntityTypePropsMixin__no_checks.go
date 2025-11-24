//go:build no_runtime_type_checking

package mixinsawsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCustomEntityTypePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCustomEntityTypePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCustomEntityTypePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCustomEntityTypePropsMixinParameters(props *CfnCustomEntityTypeMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

