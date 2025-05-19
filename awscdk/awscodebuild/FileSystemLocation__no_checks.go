//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func validateFileSystemLocation_EfsParameters(props *EfsFileSystemLocationProps) error {
	return nil
}

