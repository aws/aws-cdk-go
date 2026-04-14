//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateFileSystem_FromEfsAccessPointParameters(ap awsefs.IAccessPoint, mountPath *string) error {
	return nil
}

func validateFileSystem_FromS3FilesAccessPointParameters(ap interfacesawss3files.IAccessPointRef, mountPath *string) error {
	return nil
}

func validateNewFileSystemParameters(config *FileSystemConfig) error {
	return nil
}

