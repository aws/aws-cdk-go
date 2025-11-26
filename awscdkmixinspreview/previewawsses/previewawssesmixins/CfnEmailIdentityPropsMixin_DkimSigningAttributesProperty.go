package previewawssesmixins


// Used to configure or change the DKIM authentication settings for an email domain identity.
//
// You can use this operation to do any of the following:
//
// - Update the signing attributes for an identity that uses Bring Your Own DKIM (BYODKIM).
// - Update the key length that should be used for Easy DKIM.
// - Change from using no DKIM authentication to using Easy DKIM.
// - Change from using no DKIM authentication to using BYODKIM.
// - Change from using Easy DKIM to using BYODKIM.
// - Change from using BYODKIM to using Easy DKIM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dkimSigningAttributesProperty := &DkimSigningAttributesProperty{
//   	DomainSigningPrivateKey: jsii.String("domainSigningPrivateKey"),
//   	DomainSigningSelector: jsii.String("domainSigningSelector"),
//   	NextSigningKeyLength: jsii.String("nextSigningKeyLength"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-emailidentity-dkimsigningattributes.html
//
type CfnEmailIdentityPropsMixin_DkimSigningAttributesProperty struct {
	// [Bring Your Own DKIM] A private key that's used to generate a DKIM signature.
	//
	// The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.
	//
	// > Rather than embedding sensitive information directly in your CFN templates, we recommend you use dynamic parameters in the stack template to reference sensitive information that is stored and managed outside of CFN, such as in the AWS Systems Manager Parameter Store or AWS Secrets Manager.
	// >
	// > For more information, see the [Do not embed credentials in your templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/best-practices.html#creds) best practice.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-emailidentity-dkimsigningattributes.html#cfn-ses-emailidentity-dkimsigningattributes-domainsigningprivatekey
	//
	DomainSigningPrivateKey *string `field:"optional" json:"domainSigningPrivateKey" yaml:"domainSigningPrivateKey"`
	// [Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-emailidentity-dkimsigningattributes.html#cfn-ses-emailidentity-dkimsigningattributes-domainsigningselector
	//
	DomainSigningSelector *string `field:"optional" json:"domainSigningSelector" yaml:"domainSigningSelector"`
	// [Easy DKIM] The key length of the future DKIM key pair to be generated.
	//
	// This can be changed at most once per day.
	//
	// Valid Values: `RSA_1024_BIT | RSA_2048_BIT`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-emailidentity-dkimsigningattributes.html#cfn-ses-emailidentity-dkimsigningattributes-nextsigningkeylength
	//
	NextSigningKeyLength *string `field:"optional" json:"nextSigningKeyLength" yaml:"nextSigningKeyLength"`
}

