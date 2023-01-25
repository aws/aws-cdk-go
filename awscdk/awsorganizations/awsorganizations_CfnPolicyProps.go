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
//   cfnPolicyProps := &cfnPolicyProps{
//   	content: jsii.String("content"),
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetIds: []*string{
//   		jsii.String("targetIds"),
//   	},
//   }
//
type CfnPolicyProps struct {
	// The policy text content.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Name of the policy.
	//
	// The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) that is used to validate this parameter is a string of any of the characters in the ASCII character range.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of policy to create.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Human readable description of the policy.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of tags that you want to attach to the newly created policy.
	//
	// For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to `null` . For more information about tagging, see [Tagging AWS Organizations resources](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html) in the AWS Organizations User Guide.
	//
	// > If any one of the tags is invalid or if you exceed the allowed number of tags for a policy, then the entire request fails and the policy is not created.
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
	TargetIds *[]*string `field:"optional" json:"targetIds" yaml:"targetIds"`
}

