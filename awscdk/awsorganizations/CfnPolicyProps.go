package awsorganizations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var content interface{}
//
//   cfnPolicyProps := &CfnPolicyProps{
//   	Content: content,
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetIds: []*string{
//   		jsii.String("targetIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-policy.html
//
type CfnPolicyProps struct {
	// The policy text content. You can specify the policy content as a JSON object or a JSON string.
	//
	// > When you specify the policy content as a JSON string, you can't perform drift detection on the CloudFormation stack. For this reason, we recommend specifying the policy content as a JSON object instead.
	//
	// The text that you supply must adhere to the rules of the policy type you specify in the `Type` parameter. The following AWS Organizations quotas are enforced for the maximum size of a policy document:
	//
	// - Service control policies: 5,120 bytes *(not characters)*
	// - AI services opt-out policies: 2,500 characters
	// - Backup policies: 10,000 characters
	// - Tag policies: 10,000 characters
	//
	// For more information about Organizations service quotas, see [Quotas for AWS Organizations](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_limits.html) in the *AWS Organizations User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-policy.html#cfn-organizations-policy-content
	//
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// Name of the policy.
	//
	// The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) that is used to validate this parameter is a string of any of the characters in the ASCII character range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-policy.html#cfn-organizations-policy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of policy to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-policy.html#cfn-organizations-policy-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Human readable description of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-policy.html#cfn-organizations-policy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of tags that you want to attach to the newly created policy.
	//
	// For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to `null` . For more information about tagging, see [Tagging AWS Organizations resources](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html) in the AWS Organizations User Guide.
	//
	// > If any one of the tags is not valid or if you exceed the allowed number of tags for a policy, then the entire request fails and the policy is not created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-policy.html#cfn-organizations-policy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to.
	//
	// You can get the ID by calling the [ListRoots](https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListRoots.html) , [ListOrganizationalUnitsForParent](https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListOrganizationalUnitsForParent.html) , or [ListAccounts](https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListAccounts.html) operations. If you don't specify this parameter, the policy is created but not attached to any organization resource.
	//
	// The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) for a target ID string requires one of the following:
	//
	// - *Root* - A string that begins with "r-" followed by from 4 to 32 lowercase letters or digits.
	// - *Account* - A string that consists of exactly 12 digits.
	// - *Organizational unit (OU)* - A string that begins with "ou-" followed by from 4 to 32 lowercase letters or digits (the ID of the root that the OU is in). This string is followed by a second "-" dash and from 8 to 32 additional lowercase letters or digits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-policy.html#cfn-organizations-policy-targetids
	//
	TargetIds *[]*string `field:"optional" json:"targetIds" yaml:"targetIds"`
}

