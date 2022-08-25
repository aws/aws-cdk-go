package awsiot


// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &cfnCertificateProps{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	caCertificatePem: jsii.String("caCertificatePem"),
//   	certificateMode: jsii.String("certificateMode"),
//   	certificatePem: jsii.String("certificatePem"),
//   	certificateSigningRequest: jsii.String("certificateSigningRequest"),
//   }
//
type CfnCertificateProps struct {
	// The status of the certificate.
	//
	// Valid values are ACTIVE, INACTIVE, REVOKED, PENDING_TRANSFER, and PENDING_ACTIVATION.
	//
	// The status value REGISTER_INACTIVE is deprecated and should not be used.
	Status *string `field:"required" json:"status" yaml:"status"`
	// The CA certificate used to sign the device certificate being registered, not available when CertificateMode is SNI_ONLY.
	CaCertificatePem *string `field:"optional" json:"caCertificatePem" yaml:"caCertificatePem"`
	// Specifies which mode of certificate registration to use with this resource.
	//
	// Valid options are DEFAULT with CaCertificatePem and CertificatePem, SNI_ONLY with CertificatePem, and Default with CertificateSigningRequest.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// The certificate data in PEM format.
	//
	// Requires SNI_ONLY for the certificate mode or the accompanying CACertificatePem for registration.
	CertificatePem *string `field:"optional" json:"certificatePem" yaml:"certificatePem"`
	// The certificate signing request (CSR).
	CertificateSigningRequest *string `field:"optional" json:"certificateSigningRequest" yaml:"certificateSigningRequest"`
}

