package awsorganizations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAccount`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountProps := &CfnAccountProps{
//   	AccountName: jsii.String("accountName"),
//   	Email: jsii.String("email"),
//
//   	// the properties below are optional
//   	ParentIds: []*string{
//   		jsii.String("parentIds"),
//   	},
//   	RoleName: jsii.String("roleName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-account.html
//
type CfnAccountProps struct {
	// The account name given to the account when it was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-account.html#cfn-organizations-account-accountname
	//
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// The email address associated with the AWS account.
	//
	// The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) for this parameter is a string of characters that represents a standard internet email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-account.html#cfn-organizations-account-email
	//
	Email *string `field:"required" json:"email" yaml:"email"`
	// The unique identifier (ID) of the root or organizational unit (OU) that you want to create the new account in.
	//
	// If you don't specify this parameter, the `ParentId` defaults to the root ID.
	//
	// This parameter only accepts a string array with one string value.
	//
	// The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) for a parent ID string requires one of the following:
	//
	// - *Root* - A string that begins with "r-" followed by from 4 to 32 lowercase letters or digits.
	// - *Organizational unit (OU)* - A string that begins with "ou-" followed by from 4 to 32 lowercase letters or digits (the ID of the root that the OU is in). This string is followed by a second "-" dash and from 8 to 32 additional lowercase letters or digits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-account.html#cfn-organizations-account-parentids
	//
	ParentIds *[]*string `field:"optional" json:"parentIds" yaml:"parentIds"`
	// The name of an IAM role that AWS Organizations automatically preconfigures in the new member account.
	//
	// This role trusts the management account, allowing users in the management account to assume the role, as permitted by the management account administrator. The role has administrator permissions in the new member account.
	//
	// If you don't specify this parameter, the role name defaults to `OrganizationAccountAccessRole` .
	//
	// For more information about how to use this role to access the member account, see the following links:
	//
	// - [Accessing and Administering the Member Accounts in Your Organization](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html#orgs_manage_accounts_create-cross-account-role) in the *AWS Organizations User Guide*
	// - Steps 2 and 3 in [Tutorial: Delegate Access Across AWS accounts Using IAM Roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html) in the *IAM User Guide*
	//
	// The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) that is used to validate this parameter. The pattern can include uppercase letters, lowercase letters, digits with no spaces, and any of the following characters: =,.@-
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-account.html#cfn-organizations-account-rolename
	//
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// A list of tags that you want to attach to the newly created account.
	//
	// For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to `null` . For more information about tagging, see [Tagging AWS Organizations resources](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html) in the AWS Organizations User Guide.
	//
	// > If any one of the tags is not valid or if you exceed the maximum allowed number of tags for an account, then the entire request fails and the account is not created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-account.html#cfn-organizations-account-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

