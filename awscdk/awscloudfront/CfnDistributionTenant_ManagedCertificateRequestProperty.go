package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedCertificateRequestProperty := &ManagedCertificateRequestProperty{
//   	CertificateTransparencyLoggingPreference: jsii.String("certificateTransparencyLoggingPreference"),
//   	PrimaryDomainName: jsii.String("primaryDomainName"),
//   	ValidationTokenHost: jsii.String("validationTokenHost"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-managedcertificaterequest.html
//
type CfnDistributionTenant_ManagedCertificateRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-managedcertificaterequest.html#cfn-cloudfront-distributiontenant-managedcertificaterequest-certificatetransparencyloggingpreference
	//
	CertificateTransparencyLoggingPreference *string `field:"optional" json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-managedcertificaterequest.html#cfn-cloudfront-distributiontenant-managedcertificaterequest-primarydomainname
	//
	PrimaryDomainName *string `field:"optional" json:"primaryDomainName" yaml:"primaryDomainName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-managedcertificaterequest.html#cfn-cloudfront-distributiontenant-managedcertificaterequest-validationtokenhost
	//
	ValidationTokenHost *string `field:"optional" json:"validationTokenHost" yaml:"validationTokenHost"`
}

