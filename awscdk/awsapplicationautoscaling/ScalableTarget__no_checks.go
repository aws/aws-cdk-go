//go:build no_runtime_type_checking

package awsapplicationautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScalableTarget) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_ScalableTarget) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_ScalableTarget) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_ScalableTarget) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_ScalableTarget) validateScaleOnMetricParameters(id *string, props *BasicStepScalingPolicyProps) error {
	return nil
}

func (s *jsiiProxy_ScalableTarget) validateScaleOnScheduleParameters(id *string, action *ScalingSchedule) error {
	return nil
}

func (s *jsiiProxy_ScalableTarget) validateScaleToTrackMetricParameters(id *string, props *BasicTargetTrackingScalingPolicyProps) error {
	return nil
}

func validateScalableTarget_FromScalableTargetIdParameters(scope constructs.Construct, id *string, scalableTargetId *string) error {
	return nil
}

func validateScalableTarget_IsConstructParameters(x interface{}) error {
	return nil
}

func validateScalableTarget_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateScalableTarget_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewScalableTargetParameters(scope constructs.Construct, id *string, props *ScalableTargetProps) error {
	return nil
}

