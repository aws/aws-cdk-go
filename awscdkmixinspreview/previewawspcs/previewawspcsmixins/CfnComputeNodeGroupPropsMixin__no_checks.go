//go:build no_runtime_type_checking

package previewawspcsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnComputeNodeGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnComputeNodeGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnComputeNodeGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnComputeNodeGroupPropsMixinParameters(props *CfnComputeNodeGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

