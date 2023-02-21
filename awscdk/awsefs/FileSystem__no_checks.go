//go:build no_runtime_type_checking

package awsefs

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FileSystem) validateAddAccessPointParameters(id *string, accessPointOptions *AccessPointOptions) error {
	return nil
}

func (f *jsiiProxy_FileSystem) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FileSystem) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FileSystem) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_FileSystem) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateFileSystem_FromFileSystemAttributesParameters(scope constructs.Construct, id *string, attrs *FileSystemAttributes) error {
	return nil
}

func validateFileSystem_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFileSystem_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFileSystem_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFileSystemParameters(scope constructs.Construct, id *string, props *FileSystemProps) error {
	return nil
}

