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
//   	EnableAutoSubDomain: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html
//
type CfnDomainProps struct {
	// The unique ID for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 20.
	//
	// *Pattern:* d[a-z0-9]+.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-appid
	//
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The domain name for the domain association.
	//
	// *Length Constraints:* Maximum length of 255.
	//
	// *Pattern:* ^(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])(\.)?$
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
	// The required AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) for automatically creating subdomains.
	//
	// *Length Constraints:* Maximum length of 1000.
	//
	// *Pattern:* ^$|^arn:aws:iam::\d{12}:role.+
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-autosubdomainiamrole
	//
	AutoSubDomainIamRole *string `field:"optional" json:"autoSubDomainIamRole" yaml:"autoSubDomainIamRole"`
	// Enables the automated creation of subdomains for branches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-domain.html#cfn-amplify-domain-enableautosubdomain
	//
	EnableAutoSubDomain interface{} `field:"optional" json:"enableAutoSubDomain" yaml:"enableAutoSubDomain"`
}

