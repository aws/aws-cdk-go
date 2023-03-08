package awsiotsitewise


// Describes an asset hierarchy that contains a hierarchy's name, `LogicalID` , and child asset model ID that specifies the type of asset that can be in this hierarchy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetModelHierarchyProperty := &assetModelHierarchyProperty{
//   	childAssetModelId: jsii.String("childAssetModelId"),
//   	logicalId: jsii.String("logicalId"),
//   	name: jsii.String("name"),
//   }
//
type CfnAssetModel_AssetModelHierarchyProperty struct {
	// The Id of the asset model.
	ChildAssetModelId *string `field:"required" json:"childAssetModelId" yaml:"childAssetModelId"`
	// The `LogicalID` of the asset model hierarchy. This ID is a `hierarchyLogicalId` .
	//
	// The maximum length is 256 characters, with the pattern `[^\ u0000-\ u001F\ u007F]+`.
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The name of the asset model hierarchy.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	Name *string `field:"required" json:"name" yaml:"name"`
}

