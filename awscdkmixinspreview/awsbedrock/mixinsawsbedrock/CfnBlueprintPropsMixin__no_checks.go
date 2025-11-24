//go:build no_runtime_type_checking

package mixinsawsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBlueprintPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBlueprintPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBlueprintPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBlueprintPropsMixinParameters(props *CfnBlueprintMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

