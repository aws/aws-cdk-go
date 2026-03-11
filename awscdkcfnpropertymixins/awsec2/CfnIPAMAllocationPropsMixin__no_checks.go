//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnIPAMAllocationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnIPAMAllocationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnIPAMAllocationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnIPAMAllocationPropsMixinParameters(props *CfnIPAMAllocationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

