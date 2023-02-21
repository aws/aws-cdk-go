package awscertificatemanager


// `DomainValidationOption` is a property of the [AWS::CertificateManager::Certificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html) resource that specifies the AWS Certificate Manager ( ACM ) certificate domain to validate. Depending on the chosen validation method, ACM checks the domain's DNS record for a validation CNAME, or it attempts to send a validation email message to the domain owner.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainValidationOptionProperty := &domainValidationOptionProperty{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   	validationDomain: jsii.String("validationDomain"),
//   }
//
type CfnCertificate_DomainValidationOptionProperty struct {
	// A fully qualified domain name (FQDN) in the certificate request.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The `HostedZoneId` option, which is available if you are using Route 53 as your domain registrar, causes ACM to add your CNAME to the domain record.
	//
	// Your list of `DomainValidationOptions` must contain one and only one of the domain-validation options, and the `HostedZoneId` can be used only when `DNS` is specified as your validation method.
	//
	// Use the Route 53 `ListHostedZones` API to discover IDs for available hosted zones.
	//
	// This option is required for publicly trusted certificates.
	//
	// > The `ListHostedZones` API returns IDs in the format "/hostedzone/Z111111QQQQQQQ", but CloudFormation requires the IDs to be in the format "Z111111QQQQQQQ".
	//
	// When you change your `DomainValidationOptions` , a new resource is created.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The domain name to which you want ACM to send validation emails.
	//
	// This domain name is the suffix of the email addresses that you want ACM to use. This must be the same as the `DomainName` value or a superdomain of the `DomainName` value. For example, if you request a certificate for `testing.example.com` , you can specify `example.com` as this value. In that case, ACM sends domain validation emails to the following five addresses:
	//
	// - admin@example.com
	// - administrator@example.com
	// - hostmaster@example.com
	// - postmaster@example.com
	// - webmaster@example.com
	ValidationDomain *string `field:"optional" json:"validationDomain" yaml:"validationDomain"`
}

