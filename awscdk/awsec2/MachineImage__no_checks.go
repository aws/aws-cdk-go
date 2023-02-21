//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateMachineImage_FromSsmParameterParameters(parameterName *string, options *SsmParameterImageOptions) error {
	return nil
}

func validateMachineImage_GenericLinuxParameters(amiMap *map[string]*string, props *GenericLinuxImageProps) error {
	return nil
}

func validateMachineImage_GenericWindowsParameters(amiMap *map[string]*string, props *GenericWindowsImageProps) error {
	return nil
}

func validateMachineImage_LatestAmazonLinuxParameters(props *AmazonLinuxImageProps) error {
	return nil
}

func validateMachineImage_LatestWindowsParameters(version WindowsVersion, props *WindowsImageProps) error {
	return nil
}

func validateMachineImage_LookupParameters(props *LookupMachineImageProps) error {
	return nil
}

