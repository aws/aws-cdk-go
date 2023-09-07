package awsram

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnResourceShare`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceShareProps := &CfnResourceShareProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AllowExternalPrincipals: jsii.Boolean(false),
//   	PermissionArns: []*string{
//   		jsii.String("permissionArns"),
//   	},
//   	Principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	Sources: []*string{
//   		jsii.String("sources"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html
//
type CfnResourceShareProps struct {
	// Specifies the name of the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether principals outside your organization in AWS Organizations can be associated with a resource share.
	//
	// A value of `true` lets you share with individual AWS accounts that are *not* in your organization. A value of `false` only has meaning if your account is a member of an AWS Organization. The default value is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-allowexternalprincipals
	//
	AllowExternalPrincipals interface{} `field:"optional" json:"allowExternalPrincipals" yaml:"allowExternalPrincipals"`
	// Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) of the AWS RAM permission to associate with the resource share. If you do not specify an ARN for the permission, AWS RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-permissionarns
	//
	PermissionArns *[]*string `field:"optional" json:"permissionArns" yaml:"permissionArns"`
	// Specifies the principals to associate with the resource share. The possible values are:.
	//
	// - An AWS account ID
	// - An Amazon Resource Name (ARN) of an organization in AWS Organizations
	// - An ARN of an organizational unit (OU) in AWS Organizations
	// - An ARN of an IAM role
	// - An ARN of an IAM user
	//
	// > Not all resource types can be shared with IAM roles and users. For more information, see the column *Can share with IAM roles and users* in the tables on [Shareable AWS resources](https://docs.aws.amazon.com/ram/latest/userguide/shareable.html) in the *AWS Resource Access Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-principals
	//
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Specifies a list of one or more ARNs of the resources to associate with the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-resourcearns
	//
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-sources
	//
	Sources *[]*string `field:"optional" json:"sources" yaml:"sources"`
	// Specifies one or more tags to attach to the resource share itself.
	//
	// It doesn't attach the tags to the resources associated with the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

