//go:build no_runtime_type_checking

package mixinsawsmsk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBatchScramSecretPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBatchScramSecretPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBatchScramSecretPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBatchScramSecretPropsMixinParameters(props *CfnBatchScramSecretMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

