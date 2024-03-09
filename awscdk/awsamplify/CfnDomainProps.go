package awsamplify


// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainProps := &CfnDomainProps{
//   	AppId: jsii.String("appId"),
//   	DomainName: jsii.String("domainName"),
//   	SubDomainSettings: []interface{}{
//   		&SubDomainSettingProperty{
//   			BranchName: jsii.String("branchName"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AutoSubDomainCreationPatterns: []*string{
//   		jsii.String("autoSubDomainCreationPatterns"),
//   	},
//   	AutoSubDomainIamRole: jsii.String("autoSubDomainIamRole"),
//   	Certificate: &CertificateProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		CertificateType: jsii.String("certificateType"),
//   		CertificateVerificationDnsRecord: jsii.String("certificateVerificationDnsRecord"),
//   	},
//   	CertificateSettings: &CertificateSettingsProperty{
//   		CertificateType: jsii.String("certificateType"),
//   		CustomCertificateArn: jsii.String("customCertificateArn"),
//   	},
//   	EnableAutoSubDomain: jsii.Boolean(false),
//   	UpdateStatus: jsii.String("updateStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html
//
type CfnDomainProps struct {
	// The unique ID for an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-appid
	//
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The domain name for the domain association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The setting for the subdomain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-subdomainsettings
	//
	SubDomainSettings interface{} `field:"required" json:"subDomainSettings" yaml:"subDomainSettings"`
	// Sets the branch patterns for automatic subdomain creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-autosubdomaincreationpatterns
	//
	AutoSubDomainCreationPatterns *[]*string `field:"optional" json:"autoSubDomainCreationPatterns" yaml:"autoSubDomainCreationPatterns"`
	// The required AWS Identity and Access Management (IAMlong) service role for the Amazon Resource Name (ARN) for automatically creating subdomains.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-autosubdomainiamrole
	//
	AutoSubDomainIamRole *string `field:"optional" json:"autoSubDomainIamRole" yaml:"autoSubDomainIamRole"`
	// Describes the SSL/TLS certificate for the domain association.
	//
	// This can be your own custom certificate or the default certificate that Amplify provisions for you.
	//
	// If you are updating your domain to use a different certificate, `Certificate` points to the new certificate that is being created instead of the current active certificate. Otherwise, `Certificate` points to the current active certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-certificate
	//
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// The type of SSL/TLS certificate to use for your custom domain.
	//
	// If you don't specify a certificate type, Amplify uses the default certificate that it provisions and manages for you.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-certificatesettings
	//
	CertificateSettings interface{} `field:"optional" json:"certificateSettings" yaml:"certificateSettings"`
	// Enables the automated creation of subdomains for branches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-enableautosubdomain
	//
	EnableAutoSubDomain interface{} `field:"optional" json:"enableAutoSubDomain" yaml:"enableAutoSubDomain"`
	// The status of the domain update operation that is currently in progress.
	//
	// The following list describes the valid update states.
	//
	// - **REQUESTING_CERTIFICATE** - The certificate is in the process of being updated.
	// - **PENDING_VERIFICATION** - Indicates that an Amplify managed certificate is in the process of being verified. This occurs during the creation of a custom domain or when a custom domain is updated to use a managed certificate.
	// - **IMPORTING_CUSTOM_CERTIFICATE** - Indicates that an Amplify custom certificate is in the process of being imported. This occurs during the creation of a custom domain or when a custom domain is updated to use a custom certificate.
	// - **PENDING_DEPLOYMENT** - Indicates that the subdomain or certificate changes are being propagated.
	// - **AWAITING_APP_CNAME** - Amplify is waiting for CNAME records corresponding to subdomains to be propagated. If your custom domain is on Route 53, Amplify handles this for you automatically. For more information about custom domains, see [Setting up custom domains](https://docs.aws.amazon.com/amplify/latest/userguide/custom-domains.html) in the *Amplify Hosting User Guide* .
	// - **UPDATE_COMPLETE** - The certificate has been associated with a domain.
	// - **UPDATE_FAILED** - The certificate has failed to be provisioned or associated, and there is no existing active certificate to roll back to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-updatestatus
	//
	UpdateStatus *string `field:"optional" json:"updateStatus" yaml:"updateStatus"`
}

