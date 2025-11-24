//go:build no_runtime_type_checking

package mixinsawsefs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMountTargetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMountTargetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMountTargetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMountTargetPropsMixinParameters(props *CfnMountTargetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

