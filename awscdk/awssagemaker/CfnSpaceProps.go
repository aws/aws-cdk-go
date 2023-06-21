package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSpace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSpaceProps := &CfnSpaceProps{
//   	DomainId: jsii.String("domainId"),
//   	SpaceName: jsii.String("spaceName"),
//
//   	// the properties below are optional
//   	SpaceSettings: &SpaceSettingsProperty{
//   		JupyterServerAppSettings: &JupyterServerAppSettingsProperty{
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		KernelGatewayAppSettings: &KernelGatewayAppSettingsProperty{
//   			CustomImages: []interface{}{
//   				&CustomImageProperty{
//   					AppImageConfigName: jsii.String("appImageConfigName"),
//   					ImageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					ImageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSpaceProps struct {
	// The ID of the associated Domain.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The name of the space.
	SpaceName *string `field:"required" json:"spaceName" yaml:"spaceName"`
	// A collection of space settings.
	SpaceSettings interface{} `field:"optional" json:"spaceSettings" yaml:"spaceSettings"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

