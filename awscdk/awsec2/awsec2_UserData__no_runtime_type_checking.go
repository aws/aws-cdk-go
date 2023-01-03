//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserData) validateAddExecuteFileCommandParameters(params *ExecuteFileOptions) error {
	return nil
}

func (u *jsiiProxy_UserData) validateAddS3DownloadCommandParameters(params *S3DownloadOptions) error {
	return nil
}

func (u *jsiiProxy_UserData) validateAddSignalOnExitCommandParameters(resource awscdk.Resource) error {
	return nil
}

func validateUserData_CustomParameters(content *string) error {
	return nil
}

func validateUserData_ForLinuxParameters(options *LinuxUserDataOptions) error {
	return nil
}

func validateUserData_ForOperatingSystemParameters(os OperatingSystemType) error {
	return nil
}

func validateUserData_ForWindowsParameters(options *WindowsUserDataOptions) error {
	return nil
}

