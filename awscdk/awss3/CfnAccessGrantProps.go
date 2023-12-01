package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAccessGrant`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessGrantProps := &CfnAccessGrantProps{
//   	AccessGrantsLocationId: jsii.String("accessGrantsLocationId"),
//   	Grantee: &GranteeProperty{
//   		GranteeIdentifier: jsii.String("granteeIdentifier"),
//   		GranteeType: jsii.String("granteeType"),
//   	},
//   	Permission: jsii.String("permission"),
//
//   	// the properties below are optional
//   	AccessGrantsLocationConfiguration: &AccessGrantsLocationConfigurationProperty{
//   		S3SubPrefix: jsii.String("s3SubPrefix"),
//   	},
//   	ApplicationArn: jsii.String("applicationArn"),
//   	S3PrefixType: jsii.String("s3PrefixType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html
//
type CfnAccessGrantProps struct {
	// The ID of the registered location to which you are granting access.
	//
	// S3 Access Grants assigns this ID when you register the location. S3 Access Grants assigns the ID `default` to the default location `s3://` and assigns an auto-generated ID to other locations that you register.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-accessgrantslocationid
	//
	AccessGrantsLocationId *string `field:"required" json:"accessGrantsLocationId" yaml:"accessGrantsLocationId"`
	// The user, group, or role to which you are granting access.
	//
	// You can grant access to an IAM user or role. If you have added your corporate directory to AWS IAM Identity Center and associated your Identity Center instance with your S3 Access Grants instance, the grantee can also be a corporate directory user or group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-grantee
	//
	Grantee interface{} `field:"required" json:"grantee" yaml:"grantee"`
	// The type of access that you are granting to your S3 data, which can be set to one of the following values:  - `READ` – Grant read-only access to the S3 data.
	//
	// - `WRITE` – Grant write-only access to the S3 data.
	// - `READWRITE` – Grant both read and write access to the S3 data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-permission
	//
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// The configuration options of the grant location.
	//
	// The grant location is the S3 path to the data to which you are granting access. It contains the `S3SubPrefix` field. The grant scope is the result of appending the subprefix to the location scope of the registered location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-accessgrantslocationconfiguration
	//
	AccessGrantsLocationConfiguration interface{} `field:"optional" json:"accessGrantsLocationConfiguration" yaml:"accessGrantsLocationConfiguration"`
	// The Amazon Resource Name (ARN) of an AWS IAM Identity Center application associated with your Identity Center instance.
	//
	// If the grant includes an application ARN, the grantee can only access the S3 data through this application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-applicationarn
	//
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// The type of `S3SubPrefix` .
	//
	// The only possible value is `Object` . Pass this value if the access grant scope is an object. Do not pass this value if the access grant scope is a bucket or a bucket and a prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-s3prefixtype
	//
	S3PrefixType *string `field:"optional" json:"s3PrefixType" yaml:"s3PrefixType"`
	// The AWS resource tags that you are adding to the access grant.
	//
	// Each tag is a label consisting of a user-defined key and value. Tags can help you manage, identify, organize, search for, and filter resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

