package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAppBlockBuilder`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppBlockBuilderProps := &CfnAppBlockBuilderProps{
//   	InstanceType: jsii.String("instanceType"),
//   	Name: jsii.String("name"),
//   	Platform: jsii.String("platform"),
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AccessEndpoints: []interface{}{
//   		&AccessEndpointProperty{
//   			EndpointType: jsii.String("endpointType"),
//   			VpceId: jsii.String("vpceId"),
//   		},
//   	},
//   	AppBlockArns: []*string{
//   		jsii.String("appBlockArns"),
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	EnableDefaultInternetAccess: jsii.Boolean(false),
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html
//
type CfnAppBlockBuilderProps struct {
	// The instance type of the app block builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The name of the app block builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The platform of the app block builder.
	//
	// *Allowed values* : `WINDOWS_SERVER_2019`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-platform
	//
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// The VPC configuration for the app block builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-vpcconfig
	//
	VpcConfig interface{} `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
	// The access endpoints of the app block builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-accessendpoints
	//
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// The ARN of the app block.
	//
	// *Maximum* : `1`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-appblockarns
	//
	AppBlockArns *[]*string `field:"optional" json:"appBlockArns" yaml:"appBlockArns"`
	// The description of the app block builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the app block builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Indicates whether default internet access is enabled for the app block builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-enabledefaultinternetaccess
	//
	EnableDefaultInternetAccess interface{} `field:"optional" json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// The ARN of the IAM role that is applied to the app block builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-iamrolearn
	//
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The tags of the app block builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblockbuilder.html#cfn-appstream-appblockbuilder-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

