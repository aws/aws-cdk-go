//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) error {
	return nil
}

func (f *jsiiProxy_FlowGrants) validateStartParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FlowGrants) validateStopParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateFlowGrants_FromFlowParameters(resource interfacesawsmediaconnect.IFlowRef) error {
	return nil
}

