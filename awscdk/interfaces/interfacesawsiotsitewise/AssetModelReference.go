package interfacesawsiotsitewise


// A reference to a AssetModel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetModelReference := &AssetModelReference{
//   	AssetModelArn: jsii.String("assetModelArn"),
//   	AssetModelId: jsii.String("assetModelId"),
//   }
//
type AssetModelReference struct {
	// The ARN of the AssetModel resource.
	AssetModelArn *string `field:"required" json:"assetModelArn" yaml:"assetModelArn"`
	// The AssetModelId of the AssetModel resource.
	AssetModelId *string `field:"required" json:"assetModelId" yaml:"assetModelId"`
}

