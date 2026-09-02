//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateHlsCdnSettings_AkamaiParameters(props *HlsAkamaiCdnProps) error {
	return nil
}

func validateHlsCdnSettings_BasicPutParameters(props *HlsBasicPutCdnProps) error {
	return nil
}

func validateHlsCdnSettings_S3Parameters(props *HlsS3CdnProps) error {
	return nil
}

func validateHlsCdnSettings_WebdavParameters(props *HlsWebdavCdnProps) error {
	return nil
}

