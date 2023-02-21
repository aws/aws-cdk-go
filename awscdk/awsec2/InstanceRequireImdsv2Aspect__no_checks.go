//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceRequireImdsv2Aspect) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (i *jsiiProxy_InstanceRequireImdsv2Aspect) validateWarnParameters(node constructs.IConstruct, message *string) error {
	return nil
}

func validateNewInstanceRequireImdsv2AspectParameters(props *InstanceRequireImdsv2AspectProps) error {
	return nil
}

