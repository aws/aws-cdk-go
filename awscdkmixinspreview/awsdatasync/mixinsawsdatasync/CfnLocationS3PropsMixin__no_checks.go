//go:build no_runtime_type_checking

package mixinsawsdatasync

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLocationS3PropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLocationS3PropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLocationS3PropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLocationS3PropsMixinParameters(props *CfnLocationS3MixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

