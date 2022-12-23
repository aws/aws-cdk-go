//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_Volume) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGrantAttachVolumeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGrantAttachVolumeByResourceTagParameters(grantee awsiam.IGrantable, constructs *[]constructs.Construct) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGrantDetachVolumeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (v *jsiiProxy_Volume) validateGrantDetachVolumeByResourceTagParameters(grantee awsiam.IGrantable, constructs *[]constructs.Construct) error {
	return nil
}

func (v *jsiiProxy_Volume) validateValidatePropsParameters(props *VolumeProps) error {
	return nil
}

func validateVolume_FromVolumeAttributesParameters(scope constructs.Construct, id *string, attrs *VolumeAttributes) error {
	return nil
}

func validateVolume_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVolume_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVolume_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVolumeParameters(scope constructs.Construct, id *string, props *VolumeProps) error {
	return nil
}

