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
	// The custom S3 location to be accessed by the grantee.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-accessgrantslocationid
	//
	AccessGrantsLocationId *string `field:"required" json:"accessGrantsLocationId" yaml:"accessGrantsLocationId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-grantee
	//
	Grantee interface{} `field:"required" json:"grantee" yaml:"grantee"`
	// The level of access to be afforded to the grantee.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-permission
	//
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-accessgrantslocationconfiguration
	//
	AccessGrantsLocationConfiguration interface{} `field:"optional" json:"accessGrantsLocationConfiguration" yaml:"accessGrantsLocationConfiguration"`
	// The ARN of the application grantees will use to access the location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-applicationarn
	//
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// The type of S3SubPrefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-s3prefixtype
	//
	S3PrefixType *string `field:"optional" json:"s3PrefixType" yaml:"s3PrefixType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrant.html#cfn-s3-accessgrant-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

