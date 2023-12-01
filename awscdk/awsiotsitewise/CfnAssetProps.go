package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAsset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetProps := &CfnAssetProps{
//   	AssetModelId: jsii.String("assetModelId"),
//   	AssetName: jsii.String("assetName"),
//
//   	// the properties below are optional
//   	AssetDescription: jsii.String("assetDescription"),
//   	AssetHierarchies: []interface{}{
//   		&AssetHierarchyProperty{
//   			ChildAssetId: jsii.String("childAssetId"),
//   			LogicalId: jsii.String("logicalId"),
//   		},
//   	},
//   	AssetProperties: []interface{}{
//   		&AssetPropertyProperty{
//   			LogicalId: jsii.String("logicalId"),
//
//   			// the properties below are optional
//   			Alias: jsii.String("alias"),
//   			NotificationState: jsii.String("notificationState"),
//   			Unit: jsii.String("unit"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-asset.html
//
type CfnAssetProps struct {
	// The ID of the asset model from which to create the asset.
	//
	// This can be either the actual ID in UUID format, or else `externalId:` followed by the external ID, if it has one. For more information, see [Referencing objects with external IDs](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-id-references) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-asset.html#cfn-iotsitewise-asset-assetmodelid
	//
	AssetModelId *string `field:"required" json:"assetModelId" yaml:"assetModelId"`
	// A unique, friendly name for the asset.
	//
	// The maximum length is 256 characters with the pattern `[^\u0000-\u001F\u007F]+` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-asset.html#cfn-iotsitewise-asset-assetname
	//
	AssetName *string `field:"required" json:"assetName" yaml:"assetName"`
	// A description for the asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-asset.html#cfn-iotsitewise-asset-assetdescription
	//
	AssetDescription *string `field:"optional" json:"assetDescription" yaml:"assetDescription"`
	// A list of asset hierarchies that each contain a `hierarchyLogicalId` .
	//
	// A hierarchy specifies allowed parent/child asset relationships.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-asset.html#cfn-iotsitewise-asset-assethierarchies
	//
	AssetHierarchies interface{} `field:"optional" json:"assetHierarchies" yaml:"assetHierarchies"`
	// The list of asset properties for the asset.
	//
	// This object doesn't include properties that you define in composite models. You can find composite model properties in the `assetCompositeModels` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-asset.html#cfn-iotsitewise-asset-assetproperties
	//
	AssetProperties interface{} `field:"optional" json:"assetProperties" yaml:"assetProperties"`
	// A list of key-value pairs that contain metadata for the asset.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-asset.html#cfn-iotsitewise-asset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

