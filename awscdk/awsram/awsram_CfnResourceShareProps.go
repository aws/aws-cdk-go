package awsram

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnResourceShare`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceShareProps := &cfnResourceShareProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	allowExternalPrincipals: jsii.Boolean(false),
//   	permissionArns: []*string{
//   		jsii.String("permissionArns"),
//   	},
//   	principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	resourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnResourceShareProps struct {
	// Specifies the name of the resource share.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether principals outside your organization in AWS Organizations can be associated with a resource share.
	//
	// A value of `true` lets you share with individual AWS accounts that are *not* in your organization. A value of `false` only has meaning if your account is a member of an AWS Organization. The default value is `true` .
	AllowExternalPrincipals interface{} `field:"optional" json:"allowExternalPrincipals" yaml:"allowExternalPrincipals"`
	// Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) of the AWS RAM permission to associate with the resource share. If you do not specify an ARN for the permission, AWS RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share.
	PermissionArns *[]*string `field:"optional" json:"permissionArns" yaml:"permissionArns"`
	// Specifies a list of one or more principals to associate with the resource share.
	//
	// You can include the following values:
	//
	// - An AWS account ID, for example: `123456789012`
	// - An [Amazon Resoure Name (ARN)](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) of an organization in AWS Organizations , for example: `arn:aws:organizations::123456789012:organization/o-exampleorgid`
	// - An ARN of an organizational unit (OU) in AWS Organizations , for example: `arn:aws:organizations::123456789012:ou/o-exampleorgid/ou-examplerootid-exampleouid123`
	// - An ARN of an IAM role, for example: `arn:aws:iam::123456789012:role/rolename`
	// - An ARN of an IAM user, for example: `arn:aws:iam::123456789012user/username`
	//
	// > Not all resource types can be shared with IAM roles and users. For more information, see [Sharing with IAM roles and users](https://docs.aws.amazon.com//ram/latest/userguide/permissions.html#permissions-rbp-supported-resource-types) in the *AWS Resource Access Manager User Guide* .
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Specifies a list of one or more ARNs of the resources to associate with the resource share.
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// Specifies one or more tags to attach to the resource share itself.
	//
	// It doesn't attach the tags to the resources associated with the resource share.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

