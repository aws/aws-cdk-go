package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnUser`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnUserProps := &cfnUserProps{
//   	groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	loginProfile: &loginProfileProperty{
//   		password: jsii.String("password"),
//
//   		// the properties below are optional
//   		passwordResetRequired: jsii.Boolean(false),
//   	},
//   	managedPolicyArns: []*string{
//   		jsii.String("managedPolicyArns"),
//   	},
//   	path: jsii.String("path"),
//   	permissionsBoundary: jsii.String("permissionsBoundary"),
//   	policies: []interface{}{
//   		&policyProperty{
//   			policyDocument: policyDocument,
//   			policyName: jsii.String("policyName"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userName: jsii.String("userName"),
//   }
//
type CfnUserProps struct {
	// A list of group names to which you want to add the user.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Creates a password for the specified IAM user.
	//
	// A password allows an IAM user to access AWS services through the AWS Management Console .
	//
	// You can use the AWS CLI , the AWS API, or the *Users* page in the IAM console to create a password for any IAM user. Use [ChangePassword](https://docs.aws.amazon.com/IAM/latest/APIReference/API_ChangePassword.html) to update your own existing password in the *My Security Credentials* page in the AWS Management Console .
	//
	// For more information about managing passwords, see [Managing passwords](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html) in the *IAM User Guide* .
	LoginProfile interface{} `field:"optional" json:"loginProfile" yaml:"loginProfile"`
	// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the user.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs) and AWS Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// The path for the user name.
	//
	// For more information about paths, see [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide* .
	//
	// This parameter is optional. If it is not included, it defaults to a slash (/).
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! ( `\ u0021` ) through the DEL character ( `\ u007F` ), including most punctuation characters, digits, and upper and lowercased letters.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary *string `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// Adds or updates an inline policy document that is embedded in the specified IAM user.
	//
	// To view AWS::IAM::User snippets, see [Declaring an IAM User Resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-iam-user) .
	//
	// > The name of each policy for a role, user, or group must be unique. If you don't choose unique names, updates to the IAM identity will fail.
	//
	// For information about limits on the number of inline policies that you can embed in a user, see [Limitations on IAM Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the *IAM User Guide* .
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// A list of tags that you want to attach to the new user.
	//
	// Each tag consists of a key name and an associated value. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide* .
	//
	// > If any one of the tags is invalid or if you exceed the allowed maximum number of tags, then the entire request fails and the resource is not created.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the user to create. Do not include the path in this value.
	//
	// This parameter allows (per its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-. The user name must be unique within the account. User names are not distinguished by case. For example, you cannot create users named both "John" and "john".
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the user name.
	//
	// If you specify a name, you must specify the `CAPABILITY_NAMED_IAM` value to acknowledge your template's capabilities. For more information, see [Acknowledging IAM Resources in AWS CloudFormation Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities) .
	//
	// > Naming an IAM resource can cause an unrecoverable error if you reuse the same template in multiple Regions. To prevent this, we recommend using `Fn::Join` and `AWS::Region` to create a Region-specific name, as in the following example: `{"Fn::Join": ["", [{"Ref": "AWS::Region"}, {"Ref": "MyResourceName"}]]}` .
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

