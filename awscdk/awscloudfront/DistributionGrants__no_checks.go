//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DistributionGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) error {
	return nil
}

func (d *jsiiProxy_DistributionGrants) validateCreateInvalidationParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateDistributionGrants_FromDistributionParameters(resource interfacesawscloudfront.IDistributionRef) error {
	return nil
}

