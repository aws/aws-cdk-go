//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TargetGroupBase) validateAddLoadBalancerTargetParameters(props *LoadBalancerTargetProps) error {
	return nil
}

func (t *jsiiProxy_TargetGroupBase) validateConfigureHealthCheckParameters(healthCheck *HealthCheck) error {
	return nil
}

func (t *jsiiProxy_TargetGroupBase) validateSetAttributeParameters(key *string) error {
	return nil
}

func validateTargetGroupBase_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TargetGroupBase) validateSetHealthCheckParameters(val *HealthCheck) error {
	return nil
}

func validateNewTargetGroupBaseParameters(scope constructs.Construct, id *string, baseProps *BaseTargetGroupProps, additionalProps interface{}) error {
	return nil
}

