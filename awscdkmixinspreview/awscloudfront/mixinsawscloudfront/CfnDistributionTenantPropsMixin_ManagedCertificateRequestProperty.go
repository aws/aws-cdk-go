package mixinsawscloudfront


// An object that represents the request for the Amazon CloudFront managed ACM certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedCertificateRequestProperty := &ManagedCertificateRequestProperty{
//   	CertificateTransparencyLoggingPreference: jsii.String("certificateTransparencyLoggingPreference"),
//   	PrimaryDomainName: jsii.String("primaryDomainName"),
//   	ValidationTokenHost: jsii.String("validationTokenHost"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-managedcertificaterequest.html
//
type CfnDistributionTenantPropsMixin_ManagedCertificateRequestProperty struct {
	// You can opt out of certificate transparency logging by specifying the `disabled` option.
	//
	// Opt in by specifying `enabled` . For more information, see [Certificate Transparency Logging](https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency) in the *Certificate Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-managedcertificaterequest.html#cfn-cloudfront-distributiontenant-managedcertificaterequest-certificatetransparencyloggingpreference
	//
	CertificateTransparencyLoggingPreference *string `field:"optional" json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
	// The primary domain name associated with the CloudFront managed ACM certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-managedcertificaterequest.html#cfn-cloudfront-distributiontenant-managedcertificaterequest-primarydomainname
	//
	PrimaryDomainName *string `field:"optional" json:"primaryDomainName" yaml:"primaryDomainName"`
	// Specify how the HTTP validation token will be served when requesting the CloudFront managed ACM certificate.
	//
	// - For `cloudfront` , CloudFront will automatically serve the validation token. Choose this mode if you can point the domain's DNS to CloudFront immediately.
	// - For `self-hosted` , you serve the validation token from your existing infrastructure. Choose this mode when you need to maintain current traffic flow while your certificate is being issued. You can place the validation token at the well-known path on your existing web server, wait for ACM to validate and issue the certificate, and then update your DNS to point to CloudFront.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-managedcertificaterequest.html#cfn-cloudfront-distributiontenant-managedcertificaterequest-validationtokenhost
	//
	ValidationTokenHost *string `field:"optional" json:"validationTokenHost" yaml:"validationTokenHost"`
}

