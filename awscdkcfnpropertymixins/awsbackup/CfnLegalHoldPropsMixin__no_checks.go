//go:build no_runtime_type_checking

package awsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLegalHoldPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLegalHoldPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLegalHoldPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLegalHoldPropsMixinParameters(props *CfnLegalHoldMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

