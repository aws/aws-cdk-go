package pipelines


// An asset used by a Stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackAsset := &StackAsset{
//   	AssetId: jsii.String("assetId"),
//   	AssetManifestPath: jsii.String("assetManifestPath"),
//   	AssetSelector: jsii.String("assetSelector"),
//   	AssetType: awscdk.Pipelines.AssetType_FILE,
//   	IsTemplate: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	AssetPublishingRoleArn: jsii.String("assetPublishingRoleArn"),
//   }
//
// Experimental.
type StackAsset struct {
	// Asset identifier.
	// Experimental.
	AssetId *string `field:"required" json:"assetId" yaml:"assetId"`
	// Absolute asset manifest path.
	//
	// This needs to be made relative at a later point in time, but when this
	// information is parsed we don't know about the root cloud assembly yet.
	// Experimental.
	AssetManifestPath *string `field:"required" json:"assetManifestPath" yaml:"assetManifestPath"`
	// Asset selector to pass to `cdk-assets`.
	// Experimental.
	AssetSelector *string `field:"required" json:"assetSelector" yaml:"assetSelector"`
	// Type of asset to publish.
	// Experimental.
	AssetType AssetType `field:"required" json:"assetType" yaml:"assetType"`
	// Does this asset represent the CloudFormation template for the stack.
	// Experimental.
	IsTemplate *bool `field:"required" json:"isTemplate" yaml:"isTemplate"`
	// Role ARN to assume to publish.
	// Experimental.
	AssetPublishingRoleArn *string `field:"optional" json:"assetPublishingRoleArn" yaml:"assetPublishingRoleArn"`
}

