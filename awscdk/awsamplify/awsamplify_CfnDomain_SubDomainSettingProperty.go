package awsamplify


// The SubDomainSetting property type enables you to connect a subdomain (for example, example.exampledomain.com) to a specific branch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subDomainSettingProperty := &subDomainSettingProperty{
//   	branchName: jsii.String("branchName"),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnDomain_SubDomainSettingProperty struct {
	// The branch name setting for the subdomain.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	//
	// *Pattern:* (?s).+
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// The prefix setting for the subdomain.
	//
	// *Length Constraints:* Maximum length of 255.
	//
	// *Pattern:* (?s).*
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

