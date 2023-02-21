//go:build no_runtime_type_checking

package awsfsx

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FileSystemBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FileSystemBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FileSystemBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateFileSystemBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFileSystemBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFileSystemBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFileSystemBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

