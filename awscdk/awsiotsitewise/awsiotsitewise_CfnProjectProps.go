package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &cfnProjectProps{
//   	portalId: jsii.String("portalId"),
//   	projectName: jsii.String("projectName"),
//
//   	// the properties below are optional
//   	assetIds: []*string{
//   		jsii.String("assetIds"),
//   	},
//   	projectDescription: jsii.String("projectDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProjectProps struct {
	// The ID of the portal in which to create the project.
	PortalId *string `field:"required" json:"portalId" yaml:"portalId"`
	// A friendly name for the project.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// A list that contains the IDs of each asset associated with the project.
	AssetIds *[]*string `field:"optional" json:"assetIds" yaml:"assetIds"`
	// A description for the project.
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// A list of key-value pairs that contain metadata for the project.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

