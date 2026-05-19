//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkloadIdentity) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WorkloadIdentity) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WorkloadIdentity) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_WorkloadIdentity) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WorkloadIdentity) validateGrantAdminParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WorkloadIdentity) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WorkloadIdentity) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WorkloadIdentity) validateGrantUseParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateWorkloadIdentity_FromWorkloadIdentityAttributesParameters(scope constructs.Construct, id *string, attrs *WorkloadIdentityAttributes) error {
	return nil
}

func validateWorkloadIdentity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkloadIdentity_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWorkloadIdentity_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewWorkloadIdentityParameters(scope constructs.Construct, id *string, props *WorkloadIdentityProps) error {
	return nil
}

