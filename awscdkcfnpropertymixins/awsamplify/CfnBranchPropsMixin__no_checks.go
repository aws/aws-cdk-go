//go:build no_runtime_type_checking

package awsamplify

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBranchPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBranchPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBranchPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBranchPropsMixinParameters(props *CfnBranchMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

