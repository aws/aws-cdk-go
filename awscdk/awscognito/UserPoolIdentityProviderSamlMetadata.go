package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Metadata for a SAML user pool identity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolIdentityProviderSamlMetadata := awscdk.Aws_cognito.UserPoolIdentityProviderSamlMetadata_File(jsii.String("fileContent"))
//
type UserPoolIdentityProviderSamlMetadata interface {
	// A URL hosting SAML metadata, or the content of a file containing SAML metadata.
	MetadataContent() *string
	// The type of metadata, either a URL or file content.
	MetadataType() UserPoolIdentityProviderSamlMetadataType
}

// The jsii proxy struct for UserPoolIdentityProviderSamlMetadata
type jsiiProxy_UserPoolIdentityProviderSamlMetadata struct {
	_ byte // padding
}

func (j *jsiiProxy_UserPoolIdentityProviderSamlMetadata) MetadataContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolIdentityProviderSamlMetadata) MetadataType() UserPoolIdentityProviderSamlMetadataType {
	var returns UserPoolIdentityProviderSamlMetadataType
	_jsii_.Get(
		j,
		"metadataType",
		&returns,
	)
	return returns
}


// Specify SAML metadata via the contents of a file.
func UserPoolIdentityProviderSamlMetadata_File(fileContent *string) UserPoolIdentityProviderSamlMetadata {
	_init_.Initialize()

	if err := validateUserPoolIdentityProviderSamlMetadata_FileParameters(fileContent); err != nil {
		panic(err)
	}
	var returns UserPoolIdentityProviderSamlMetadata

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolIdentityProviderSamlMetadata",
		"file",
		[]interface{}{fileContent},
		&returns,
	)

	return returns
}

// Specify SAML metadata via a URL.
func UserPoolIdentityProviderSamlMetadata_Url(url *string) UserPoolIdentityProviderSamlMetadata {
	_init_.Initialize()

	if err := validateUserPoolIdentityProviderSamlMetadata_UrlParameters(url); err != nil {
		panic(err)
	}
	var returns UserPoolIdentityProviderSamlMetadata

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolIdentityProviderSamlMetadata",
		"url",
		[]interface{}{url},
		&returns,
	)

	return returns
}

