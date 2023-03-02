package awsiotsitewise


// Describes an asset hierarchy that contains a `childAssetId` and `hierarchyLogicalId` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetHierarchyProperty := &assetHierarchyProperty{
//   	childAssetId: jsii.String("childAssetId"),
//   	logicalId: jsii.String("logicalId"),
//   }
//
type CfnAsset_AssetHierarchyProperty struct {
	// The Id of the child asset.
	ChildAssetId *string `field:"required" json:"childAssetId" yaml:"childAssetId"`
	// The `LogicalID` of the hierarchy. This ID is a `hierarchyLogicalId` .
	//
	// The maximum length is 256 characters, with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
}

