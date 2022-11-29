package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAsset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetProps := &cfnAssetProps{
//   	assetModelId: jsii.String("assetModelId"),
//   	assetName: jsii.String("assetName"),
//
//   	// the properties below are optional
//   	assetDescription: jsii.String("assetDescription"),
//   	assetHierarchies: []interface{}{
//   		&assetHierarchyProperty{
//   			childAssetId: jsii.String("childAssetId"),
//   			logicalId: jsii.String("logicalId"),
//   		},
//   	},
//   	assetProperties: []interface{}{
//   		&assetPropertyProperty{
//   			logicalId: jsii.String("logicalId"),
//
//   			// the properties below are optional
//   			alias: jsii.String("alias"),
//   			notificationState: jsii.String("notificationState"),
//   			unit: jsii.String("unit"),
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
type CfnAssetProps struct {
	// The ID of the asset model from which to create the asset.
	AssetModelId *string `field:"required" json:"assetModelId" yaml:"assetModelId"`
	// A unique, friendly name for the asset.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	AssetName *string `field:"required" json:"assetName" yaml:"assetName"`
	// `AWS::IoTSiteWise::Asset.AssetDescription`.
	AssetDescription *string `field:"optional" json:"assetDescription" yaml:"assetDescription"`
	// A list of asset hierarchies that each contain a `hierarchyLogicalId` .
	//
	// A hierarchy specifies allowed parent/child asset relationships.
	AssetHierarchies interface{} `field:"optional" json:"assetHierarchies" yaml:"assetHierarchies"`
	// The list of asset properties for the asset.
	//
	// This object doesn't include properties that you define in composite models. You can find composite model properties in the `assetCompositeModels` object.
	AssetProperties interface{} `field:"optional" json:"assetProperties" yaml:"assetProperties"`
	// A list of key-value pairs that contain metadata for the asset.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

