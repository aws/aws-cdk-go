package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSpace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSpaceProps := &cfnSpaceProps{
//   	domainId: jsii.String("domainId"),
//   	spaceName: jsii.String("spaceName"),
//
//   	// the properties below are optional
//   	spaceSettings: &spaceSettingsProperty{
//   		jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   			customImages: []interface{}{
//   				&customImageProperty{
//   					appImageConfigName: jsii.String("appImageConfigName"),
//   					imageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					imageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

