package pipelines


// Instructions to publish certain assets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPublishingCommand := &assetPublishingCommand{
//   	assetId: jsii.String("assetId"),
//   	assetManifestPath: jsii.String("assetManifestPath"),
//   	assetPublishingRoleArn: jsii.String("assetPublishingRoleArn"),
//   	assetSelector: jsii.String("assetSelector"),
//   	assetType: awscdk.Pipelines.assetType_FILE,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type AssetPublishingCommand struct {
	// Asset identifier.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetId *string `field:"required" json:"assetId" yaml:"assetId"`
	// Asset manifest path.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetManifestPath *string `field:"required" json:"assetManifestPath" yaml:"assetManifestPath"`
	// ARN of the IAM Role used to publish this asset.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetPublishingRoleArn *string `field:"required" json:"assetPublishingRoleArn" yaml:"assetPublishingRoleArn"`
	// Asset selector to pass to `cdk-assets`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetSelector *string `field:"required" json:"assetSelector" yaml:"assetSelector"`
	// Type of asset to publish.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetType AssetType `field:"required" json:"assetType" yaml:"assetType"`
}

