//go:build no_runtime_type_checking

package awssecretsmanager

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RotationSchedule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RotationSchedule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RotationSchedule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_RotationSchedule) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (r *jsiiProxy_RotationSchedule) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateRotationSchedule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRotationSchedule_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewRotationScheduleParameters(scope constructs.Construct, id *string, props *RotationScheduleProps) error {
	return nil
}

