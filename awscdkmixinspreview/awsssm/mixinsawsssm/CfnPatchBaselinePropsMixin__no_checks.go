//go:build no_runtime_type_checking

package mixinsawsssm

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPatchBaselinePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPatchBaselinePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPatchBaselinePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPatchBaselinePropsMixinParameters(props *CfnPatchBaselineMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

