package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A SAML metadata document.
//
// Example:
//   provider := iam.NewSamlProvider(this, jsii.String("Provider"), &SamlProviderProps{
//   	MetadataDocument: iam.SamlMetadataDocument_FromFile(jsii.String("/path/to/saml-metadata-document.xml")),
//   })
//   principal := iam.NewSamlPrincipal(provider, map[string]interface{}{
//   	"StringEquals": map[string]*string{
//   		"SAML:iss": jsii.String("issuer"),
//   	},
//   })
//
type SamlMetadataDocument interface {
	// The XML content of the metadata document.
	Xml() *string
}

// The jsii proxy struct for SamlMetadataDocument
type jsiiProxy_SamlMetadataDocument struct {
	_ byte // padding
}

func (j *jsiiProxy_SamlMetadataDocument) Xml() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xml",
		&returns,
	)
	return returns
}


func NewSamlMetadataDocument_Override(s SamlMetadataDocument) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.SamlMetadataDocument",
		nil, // no parameters
		s,
	)
}

// Create a SAML metadata document from a XML file.
func SamlMetadataDocument_FromFile(path *string) SamlMetadataDocument {
	_init_.Initialize()

	if err := validateSamlMetadataDocument_FromFileParameters(path); err != nil {
		panic(err)
	}
	var returns SamlMetadataDocument

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.SamlMetadataDocument",
		"fromFile",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Create a SAML metadata document from a XML string.
func SamlMetadataDocument_FromXml(xml *string) SamlMetadataDocument {
	_init_.Initialize()

	if err := validateSamlMetadataDocument_FromXmlParameters(xml); err != nil {
		panic(err)
	}
	var returns SamlMetadataDocument

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.SamlMetadataDocument",
		"fromXml",
		[]interface{}{xml},
		&returns,
	)

	return returns
}

