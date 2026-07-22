package interfacesawsquicksight


// A reference to a AssetBundleImportJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetBundleImportJobReference := &AssetBundleImportJobReference{
//   	AssetBundleImportJobArn: jsii.String("assetBundleImportJobArn"),
//   	AssetBundleImportJobId: jsii.String("assetBundleImportJobId"),
//   }
//
type AssetBundleImportJobReference struct {
	// The ARN of the AssetBundleImportJob resource.
	AssetBundleImportJobArn *string `field:"required" json:"assetBundleImportJobArn" yaml:"assetBundleImportJobArn"`
	// The AssetBundleImportJobId of the AssetBundleImportJob resource.
	AssetBundleImportJobId *string `field:"required" json:"assetBundleImportJobId" yaml:"assetBundleImportJobId"`
}

