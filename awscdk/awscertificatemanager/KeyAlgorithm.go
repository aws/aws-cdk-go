package awscertificatemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Certificate Manager key algorithm.
//
// If you need to use an algorithm that doesn't exist as a static member, you
// can instantiate a `KeyAlgorithm` object, e.g: `new KeyAlgorithm('RSA_2048')`.
//
// Example:
//   import acmpca "github.com/aws/aws-cdk-go/awscdk"
//
//
//   acm.NewPrivateCertificate(this, jsii.String("PrivateCertificate"), &PrivateCertificateProps{
//   	DomainName: jsii.String("test.example.com"),
//   	SubjectAlternativeNames: []*string{
//   		jsii.String("cool.example.com"),
//   		jsii.String("test.example.net"),
//   	},
//   	 // optional
//   	CertificateAuthority: acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("CA"), jsii.String("arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/023077d8-2bfa-4eb0-8f22-05c96deade77")),
//   	KeyAlgorithm: acm.KeyAlgorithm_RSA_2048(),
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-keyalgorithm
//
type KeyAlgorithm interface {
	// The name of the algorithm.
	Name() *string
}

// The jsii proxy struct for KeyAlgorithm
type jsiiProxy_KeyAlgorithm struct {
	_ byte // padding
}

func (j *jsiiProxy_KeyAlgorithm) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewKeyAlgorithm(name *string) KeyAlgorithm {
	_init_.Initialize()

	if err := validateNewKeyAlgorithmParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyAlgorithm{}

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.KeyAlgorithm",
		[]interface{}{name},
		&j,
	)

	return &j
}

func NewKeyAlgorithm_Override(k KeyAlgorithm, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_certificatemanager.KeyAlgorithm",
		[]interface{}{name},
		k,
	)
}

func KeyAlgorithm_EC_PRIME256V1() KeyAlgorithm {
	_init_.Initialize()
	var returns KeyAlgorithm
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_certificatemanager.KeyAlgorithm",
		"EC_PRIME256V1",
		&returns,
	)
	return returns
}

func KeyAlgorithm_EC_SECP384R1() KeyAlgorithm {
	_init_.Initialize()
	var returns KeyAlgorithm
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_certificatemanager.KeyAlgorithm",
		"EC_SECP384R1",
		&returns,
	)
	return returns
}

func KeyAlgorithm_RSA_2048() KeyAlgorithm {
	_init_.Initialize()
	var returns KeyAlgorithm
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_certificatemanager.KeyAlgorithm",
		"RSA_2048",
		&returns,
	)
	return returns
}

