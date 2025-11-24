package mixinsawsamplify


// The SubDomainSetting property type enables you to connect a subdomain (for example, example.exampledomain.com) to a specific branch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subDomainSettingProperty := &SubDomainSettingProperty{
//   	BranchName: jsii.String("branchName"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-subdomainsetting.html
//
type CfnDomainPropsMixin_SubDomainSettingProperty struct {
	// The branch name setting for the subdomain.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	//
	// *Pattern:* (?s).+
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-subdomainsetting.html#cfn-amplify-domain-subdomainsetting-branchname
	//
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// The prefix setting for the subdomain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-subdomainsetting.html#cfn-amplify-domain-subdomainsetting-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

