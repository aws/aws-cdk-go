//go:build no_runtime_type_checking

package awscdklocationalpha

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Tracker) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_Tracker) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_Tracker) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_Tracker) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Tracker) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Tracker) validateGrantUpdateDevicePositionsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateTracker_FromTrackerArnParameters(scope constructs.Construct, id *string, trackerArn *string) error {
	return nil
}

func validateTracker_FromTrackerNameParameters(scope constructs.Construct, id *string, trackerName *string) error {
	return nil
}

func validateTracker_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTracker_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTracker_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTrackerParameters(scope constructs.Construct, id *string, props *TrackerProps) error {
	return nil
}

