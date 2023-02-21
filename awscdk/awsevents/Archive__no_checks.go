//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Archive) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Archive) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Archive) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateArchive_IsConstructParameters(x interface{}) error {
	return nil
}

func validateArchive_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateArchive_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewArchiveParameters(scope constructs.Construct, id *string, props *ArchiveProps) error {
	return nil
}

