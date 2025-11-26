package previewawsamplifymixins


// Properties for CfnDomainPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDomainMixinProps := &CfnDomainMixinProps{
//   	AppId: jsii.String("appId"),
//   	AutoSubDomainCreationPatterns: []*string{
//   		jsii.String("autoSubDomainCreationPatterns"),
//   	},
//   	AutoSubDomainIamRole: jsii.String("autoSubDomainIamRole"),
//   	CertificateSettings: &CertificateSettingsProperty{
//   		CertificateType: jsii.String("certificateType"),
//   		CustomCertificateArn: jsii.String("customCertificateArn"),
//   	},
//   	DomainName: jsii.String("domainName"),
//   	EnableAutoSubDomain: jsii.Boolean(false),
//   	SubDomainSettings: []interface{}{
//   		&SubDomainSettingProperty{
//   			BranchName: jsii.String("branchName"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html
//
type CfnDomainMixinProps struct {
	// The unique ID for an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-appid
	//
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
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
	// The domain name for the domain association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Enables the automated creation of subdomains for branches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-enableautosubdomain
	//
	EnableAutoSubDomain interface{} `field:"optional" json:"enableAutoSubDomain" yaml:"enableAutoSubDomain"`
	// The setting for the subdomain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-subdomainsettings
	//
	SubDomainSettings interface{} `field:"optional" json:"subDomainSettings" yaml:"subDomainSettings"`
}

