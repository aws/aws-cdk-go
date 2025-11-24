//go:build no_runtime_type_checking

package mixinsawsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnComputeEnvironmentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnComputeEnvironmentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnComputeEnvironmentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnComputeEnvironmentPropsMixinParameters(props *CfnComputeEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

