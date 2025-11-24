//go:build no_runtime_type_checking

package mixinsawslightsail

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDiskPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDiskPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDiskPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDiskPropsMixinParameters(props *CfnDiskMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

