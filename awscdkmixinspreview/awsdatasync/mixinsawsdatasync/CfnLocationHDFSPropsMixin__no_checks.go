//go:build no_runtime_type_checking

package mixinsawsdatasync

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLocationHDFSPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLocationHDFSPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLocationHDFSPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLocationHDFSPropsMixinParameters(props *CfnLocationHDFSMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

