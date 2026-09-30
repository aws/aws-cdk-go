//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSubnetCidrReservationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSubnetCidrReservationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSubnetCidrReservationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSubnetCidrReservationPropsMixinParameters(props *CfnSubnetCidrReservationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

