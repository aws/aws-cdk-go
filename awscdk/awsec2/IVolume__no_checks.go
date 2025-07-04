//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVolume) validateGrantAttachVolumeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IVolume) validateGrantAttachVolumeByResourceTagParameters(grantee awsiam.IGrantable, constructs *[]constructs.Construct) error {
	return nil
}

func (i *jsiiProxy_IVolume) validateGrantDetachVolumeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IVolume) validateGrantDetachVolumeByResourceTagParameters(grantee awsiam.IGrantable, constructs *[]constructs.Construct) error {
	return nil
}

