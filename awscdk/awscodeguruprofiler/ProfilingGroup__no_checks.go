//go:build no_runtime_type_checking

package awscodeguruprofiler

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProfilingGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_ProfilingGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_ProfilingGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_ProfilingGroup) validateGrantPublishParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_ProfilingGroup) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateProfilingGroup_FromProfilingGroupArnParameters(scope constructs.Construct, id *string, profilingGroupArn *string) error {
	return nil
}

func validateProfilingGroup_FromProfilingGroupNameParameters(scope constructs.Construct, id *string, profilingGroupName *string) error {
	return nil
}

func validateProfilingGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProfilingGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateProfilingGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewProfilingGroupParameters(scope constructs.Construct, id *string, props *ProfilingGroupProps) error {
	return nil
}

