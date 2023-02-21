//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateFileSystem_CopyDirectoryParameters(srcDir *string, destDir *string, options *CopyOptions) error {
	return nil
}

func validateFileSystem_FingerprintParameters(fileOrDirectory *string, options *FingerprintOptions) error {
	return nil
}

func validateFileSystem_IsEmptyParameters(dir *string) error {
	return nil
}

func validateFileSystem_MkdtempParameters(prefix *string) error {
	return nil
}

