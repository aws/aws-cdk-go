//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MultipartUserData) validateAddExecuteFileCommandParameters(params *ExecuteFileOptions) error {
	return nil
}

func (m *jsiiProxy_MultipartUserData) validateAddPartParameters(part MultipartBody) error {
	return nil
}

func (m *jsiiProxy_MultipartUserData) validateAddS3DownloadCommandParameters(params *S3DownloadOptions) error {
	return nil
}

func (m *jsiiProxy_MultipartUserData) validateAddSignalOnExitCommandParameters(resource awscdk.Resource) error {
	return nil
}

func (m *jsiiProxy_MultipartUserData) validateAddUserDataPartParameters(userData UserData) error {
	return nil
}

func validateMultipartUserData_CustomParameters(content *string) error {
	return nil
}

func validateMultipartUserData_ForLinuxParameters(options *LinuxUserDataOptions) error {
	return nil
}

func validateMultipartUserData_ForOperatingSystemParameters(os OperatingSystemType) error {
	return nil
}

func validateMultipartUserData_ForWindowsParameters(options *WindowsUserDataOptions) error {
	return nil
}

func validateNewMultipartUserDataParameters(opts *MultipartUserDataOptions) error {
	return nil
}

