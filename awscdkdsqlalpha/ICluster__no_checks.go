//go:build no_runtime_type_checking

package awscdkdsqlalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICluster) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ICluster) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ICluster) validateGrantConnectAdminParameters(grantee awsiam.IGrantable) error {
	return nil
}

