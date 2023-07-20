package awsiotsitewise


// Describes an asset hierarchy that contains a `childAssetId` and `hierarchyLogicalId` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetHierarchyProperty := &AssetHierarchyProperty{
//   	ChildAssetId: jsii.String("childAssetId"),
//   	LogicalId: jsii.String("logicalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assethierarchy.html
//
type CfnAsset_AssetHierarchyProperty struct {
	// The Id of the child asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assethierarchy.html#cfn-iotsitewise-asset-assethierarchy-childassetid
	//
	ChildAssetId *string `field:"required" json:"childAssetId" yaml:"childAssetId"`
	// The `LogicalID` of the hierarchy. This ID is a `hierarchyLogicalId` .
	//
	// The maximum length is 256 characters, with the pattern `[^\u0000-\u001F\u007F]+` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assethierarchy.html#cfn-iotsitewise-asset-assethierarchy-logicalid
	//
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
}

