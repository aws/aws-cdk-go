//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CrossRegionInferenceProfile) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CrossRegionInferenceProfile) validateGrantProfileUsageParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateCrossRegionInferenceProfile_FromConfigParameters(config *CrossRegionInferenceProfileProps) error {
	return nil
}

