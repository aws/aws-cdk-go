//go:build no_runtime_type_checking

package awsecr

// Building without runtime type checking enabled, so all the below just return nil

func validatePublicGalleryAuthorizationToken_GrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

