//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateInitPackage_AptParameters(packageName *string, options *NamedPackageOptions) error {
	return nil
}

func validateInitPackage_MsiParameters(location *string, options *LocationPackageOptions) error {
	return nil
}

func validateInitPackage_PythonParameters(packageName *string, options *NamedPackageOptions) error {
	return nil
}

func validateInitPackage_RpmParameters(location *string, options *LocationPackageOptions) error {
	return nil
}

func validateInitPackage_RubyGemParameters(gemName *string, options *NamedPackageOptions) error {
	return nil
}

func validateInitPackage_YumParameters(packageName *string, options *NamedPackageOptions) error {
	return nil
}

func validateNewInitPackageParameters(type_ *string, versions *[]*string) error {
	return nil
}

