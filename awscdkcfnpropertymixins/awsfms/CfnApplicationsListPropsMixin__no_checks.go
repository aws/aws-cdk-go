//go:build no_runtime_type_checking

package awsfms

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnApplicationsListPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationsListPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnApplicationsListPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnApplicationsListPropsMixinParameters(props *CfnApplicationsListMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

