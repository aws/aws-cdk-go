package pipelines


// An asset used by a Stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackAsset := &stackAsset{
//   	assetId: jsii.String("assetId"),
//   	assetManifestPath: jsii.String("assetManifestPath"),
//   	assetSelector: jsii.String("assetSelector"),
//   	assetType: awscdk.Pipelines.assetType_FILE,
//   	isTemplate: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	assetPublishingRoleArn: jsii.String("assetPublishingRoleArn"),
//   }
//
type StackAsset struct {
	// Asset identifier.
	AssetId *string `field:"required" json:"assetId" yaml:"assetId"`
	// Absolute asset manifest path.
	//
	// This needs to be made relative at a later point in time, but when this
	// information is parsed we don't know about the root cloud assembly yet.
	AssetManifestPath *string `field:"required" json:"assetManifestPath" yaml:"assetManifestPath"`
	// Asset selector to pass to `cdk-assets`.
	AssetSelector *string `field:"required" json:"assetSelector" yaml:"assetSelector"`
	// Type of asset to publish.
	AssetType AssetType `field:"required" json:"assetType" yaml:"assetType"`
	// Does this asset represent the CloudFormation template for the stack.
	IsTemplate *bool `field:"required" json:"isTemplate" yaml:"isTemplate"`
	// Role ARN to assume to publish.
	AssetPublishingRoleArn *string `field:"optional" json:"assetPublishingRoleArn" yaml:"assetPublishingRoleArn"`
}

