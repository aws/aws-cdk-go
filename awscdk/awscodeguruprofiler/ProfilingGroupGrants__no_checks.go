//go:build no_runtime_type_checking

package awscodeguruprofiler

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProfilingGroupGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) error {
	return nil
}

func (p *jsiiProxy_ProfilingGroupGrants) validatePublishParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_ProfilingGroupGrants) validateReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateProfilingGroupGrants_FromProfilingGroupParameters(resource interfacesawscodeguruprofiler.IProfilingGroupRef) error {
	return nil
}

