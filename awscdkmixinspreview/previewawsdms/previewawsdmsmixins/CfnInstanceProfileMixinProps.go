package previewawsdmsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnInstanceProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceProfileMixinProps := &CfnInstanceProfileMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Description: jsii.String("description"),
//   	InstanceProfileIdentifier: jsii.String("instanceProfileIdentifier"),
//   	InstanceProfileName: jsii.String("instanceProfileName"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	NetworkType: jsii.String("networkType"),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	SubnetGroupIdentifier: jsii.String("subnetGroupIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcSecurityGroups: []*string{
//   		jsii.String("vpcSecurityGroups"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html
//
type CfnInstanceProfileMixinProps struct {
	// The Availability Zone where the instance profile runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html#cfn-dms-instanceprofile-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// A description of the instance profile.
	//
	// Descriptions can have up to 31 characters. A description can contain only ASCII letters, digits, and hyphens ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html#cfn-dms-instanceprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of the instance profile.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html#cfn-dms-instanceprofile-instanceprofileidentifier
	//
	InstanceProfileIdentifier *string `field:"optional" json:"instanceProfileIdentifier" yaml:"instanceProfileIdentifier"`
	// The user-friendly name for the instance profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html#cfn-dms-instanceprofile-instanceprofilename
	//
	InstanceProfileName *string `field:"optional" json:"instanceProfileName" yaml:"instanceProfileName"`
	// The Amazon Resource Name (ARN) of the AWS  key that is used to encrypt the connection parameters for the instance profile.
	//
	// If you don't specify a value for the `KmsKeyArn` parameter, then AWS DMS uses an AWS owned encryption key to encrypt your resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html#cfn-dms-instanceprofile-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Specifies the network type for the instance profile.
	//
	// A value of `IPV4` represents an instance profile with IPv4 network type and only supports IPv4 addressing. A value of `IPV6` represents an instance profile with IPv6 network type and only supports IPv6 addressing. A value of `DUAL` represents an instance profile with dual network type that supports IPv4 and IPv6 addressing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html#cfn-dms-instanceprofile-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Specifies the accessibility options for the instance profile.
	//
	// A value of `true` represents an instance profile with a public IP address. A value of `false` represents an instance profile with a private IP address. The default value is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html#cfn-dms-instanceprofile-publiclyaccessible
	//
	// Default: - false.
	//
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The identifier of the subnet group that is associated with the instance profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html#cfn-dms-instanceprofile-subnetgroupidentifier
	//
	SubnetGroupIdentifier *string `field:"optional" json:"subnetGroupIdentifier" yaml:"subnetGroupIdentifier"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html#cfn-dms-instanceprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The VPC security groups that are used with the instance profile.
	//
	// The VPC security group must work with the VPC containing the instance profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-instanceprofile.html#cfn-dms-instanceprofile-vpcsecuritygroups
	//
	VpcSecurityGroups *[]*string `field:"optional" json:"vpcSecurityGroups" yaml:"vpcSecurityGroups"`
}

