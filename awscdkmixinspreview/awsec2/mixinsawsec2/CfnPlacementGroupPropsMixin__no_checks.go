//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPlacementGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPlacementGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPlacementGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPlacementGroupPropsMixinParameters(props *CfnPlacementGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

