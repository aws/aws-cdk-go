//go:build no_runtime_type_checking

package awsfsx

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LustreFileSystem) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LustreFileSystem) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LustreFileSystem) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateLustreFileSystem_FromLustreFileSystemAttributesParameters(scope constructs.Construct, id *string, attrs *FileSystemAttributes) error {
	return nil
}

func validateLustreFileSystem_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLustreFileSystem_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLustreFileSystem_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLustreFileSystemParameters(scope constructs.Construct, id *string, props *LustreFileSystemProps) error {
	return nil
}

