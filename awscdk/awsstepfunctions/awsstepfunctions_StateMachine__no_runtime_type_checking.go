//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StateMachine) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateGrantParameters(identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateGrantExecutionParameters(identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateGrantReadParameters(identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateGrantStartExecutionParameters(identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateGrantStartSyncExecutionParameters(identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateGrantTaskResponseParameters(identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateMetricAbortedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateMetricThrottledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StateMachine) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateStateMachine_FromStateMachineArnParameters(scope constructs.Construct, id *string, stateMachineArn *string) error {
	return nil
}

func validateStateMachine_FromStateMachineNameParameters(scope constructs.Construct, id *string, stateMachineName *string) error {
	return nil
}

func validateStateMachine_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStateMachine_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateStateMachine_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewStateMachineParameters(scope constructs.Construct, id *string, props *StateMachineProps) error {
	return nil
}

