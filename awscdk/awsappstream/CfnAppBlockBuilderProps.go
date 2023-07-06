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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAppBlockBuilderProps struct {
	// The instance type of the app block builder.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The name of the app block builder.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The platform of the app block builder.
	//
	// *Allowed values* : `WINDOWS_SERVER_2019`.
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// The VPC configuration for the app block builder.
	VpcConfig interface{} `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
	// The access endpoints of the app block builder.
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// The ARN of the app block.
	//
	// *Maximum* : `1`.
	AppBlockArns *[]*string `field:"optional" json:"appBlockArns" yaml:"appBlockArns"`
	// The description of the app block builder.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the app block builder.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Indicates whether default internet access is enabled for the app block builder.
	EnableDefaultInternetAccess interface{} `field:"optional" json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// The ARN of the IAM role that is applied to the app block builder.
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The tags of the app block builder.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

