//go:build no_runtime_type_checking

package awscodeguruprofiler

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IProfilingGroup) validateGrantPublishParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IProfilingGroup) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

