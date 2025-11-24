//go:build no_runtime_type_checking

package mixinsawstimestream

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInfluxDBInstancePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInfluxDBInstancePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInfluxDBInstancePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInfluxDBInstancePropsMixinParameters(props *CfnInfluxDBInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

