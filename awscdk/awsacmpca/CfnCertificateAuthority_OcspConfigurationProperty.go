package awsacmpca


// Contains information to enable and configure Online Certificate Status Protocol (OCSP) for validating certificate revocation status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ocspConfigurationProperty := &OcspConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	OcspCustomCname: jsii.String("ocspCustomCname"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-ocspconfiguration.html
//
type CfnCertificateAuthority_OcspConfigurationProperty struct {
	// Flag enabling use of the Online Certificate Status Protocol (OCSP) for validating certificate revocation status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-ocspconfiguration.html#cfn-acmpca-certificateauthority-ocspconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// By default, AWS Private CA injects an Amazon domain into certificates being validated by the Online Certificate Status Protocol (OCSP).
	//
	// A customer can alternatively use this object to define a CNAME specifying a customized OCSP domain.
	//
	// > The content of a Canonical Name (CNAME) record must conform to [RFC2396](https://docs.aws.amazon.com/https://www.ietf.org/rfc/rfc2396.txt) restrictions on the use of special characters in URIs. Additionally, the value of the CNAME must not include a protocol prefix such as "http://" or "https://".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-ocspconfiguration.html#cfn-acmpca-certificateauthority-ocspconfiguration-ocspcustomcname
	//
	OcspCustomCname *string `field:"optional" json:"ocspCustomCname" yaml:"ocspCustomCname"`
}

