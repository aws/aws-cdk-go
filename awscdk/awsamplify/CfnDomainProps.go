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
//   	CertificateSettings: &CertificateSettingsProperty{
//   		CertificateType: jsii.String("certificateType"),
//   		CustomCertificateArn: jsii.String("customCertificateArn"),
//   	},
//   	EnableAutoSubDomain: jsii.Boolean(false),
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
}

