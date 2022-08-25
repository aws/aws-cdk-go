package awsacmpca


// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.
//
// Your private certificate authority (CA) can configure Online Certificate Status Protocol (OCSP) support and/or maintain a certificate revocation list (CRL). OCSP returns validation information about certificates as requested by clients, and a CRL contains an updated list of certificates revoked by your CA. For more information, see [RevokeCertificate](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_RevokeCertificate.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   revocationConfigurationProperty := &revocationConfigurationProperty{
//   	crlConfiguration: &crlConfigurationProperty{
//   		customCname: jsii.String("customCname"),
//   		enabled: jsii.Boolean(false),
//   		expirationInDays: jsii.Number(123),
//   		s3BucketName: jsii.String("s3BucketName"),
//   		s3ObjectAcl: jsii.String("s3ObjectAcl"),
//   	},
//   	ocspConfiguration: &ocspConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   		ocspCustomCname: jsii.String("ocspCustomCname"),
//   	},
//   }
//
type CfnCertificateAuthority_RevocationConfigurationProperty struct {
	// Configuration of the certificate revocation list (CRL), if any, maintained by your private CA.
	CrlConfiguration interface{} `field:"optional" json:"crlConfiguration" yaml:"crlConfiguration"`
	// Configuration of Online Certificate Status Protocol (OCSP) support, if any, maintained by your private CA.
	OcspConfiguration interface{} `field:"optional" json:"ocspConfiguration" yaml:"ocspConfiguration"`
}

