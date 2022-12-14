package awsacmpca


// Contains information to enable and configure Online Certificate Status Protocol (OCSP) for validating certificate revocation status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ocspConfigurationProperty := &ocspConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   	ocspCustomCname: jsii.String("ocspCustomCname"),
//   }
//
type CfnCertificateAuthority_OcspConfigurationProperty struct {
	// Flag enabling use of the Online Certificate Status Protocol (OCSP) for validating certificate revocation status.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// By default, ACM Private CA injects an Amazon domain into certificates being validated by the Online Certificate Status Protocol (OCSP).
	//
	// A customer can alternatively use this object to define a CNAME specifying a customized OCSP domain.
	//
	// Note: The value of the CNAME must not include a protocol prefix such as "http://" or "https://".
	OcspCustomCname *string `field:"optional" json:"ocspCustomCname" yaml:"ocspCustomCname"`
}

