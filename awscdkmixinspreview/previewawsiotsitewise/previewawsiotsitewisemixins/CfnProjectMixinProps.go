package previewawsiotsitewisemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnProjectPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProjectMixinProps := &CfnProjectMixinProps{
//   	AssetIds: []*string{
//   		jsii.String("assetIds"),
//   	},
//   	PortalId: jsii.String("portalId"),
//   	ProjectDescription: jsii.String("projectDescription"),
//   	ProjectName: jsii.String("projectName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-project.html
//
type CfnProjectMixinProps struct {
	// A list that contains the IDs of each asset associated with the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-project.html#cfn-iotsitewise-project-assetids
	//
	AssetIds *[]*string `field:"optional" json:"assetIds" yaml:"assetIds"`
	// The ID of the portal in which to create the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-project.html#cfn-iotsitewise-project-portalid
	//
	PortalId *string `field:"optional" json:"portalId" yaml:"portalId"`
	// A description for the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-project.html#cfn-iotsitewise-project-projectdescription
	//
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// A friendly name for the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-project.html#cfn-iotsitewise-project-projectname
	//
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// A list of key-value pairs that contain metadata for the project.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-project.html#cfn-iotsitewise-project-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

