package interfacesawsquicksight


// A reference to a AssetBundleExportJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetBundleExportJobReference := &AssetBundleExportJobReference{
//   	AssetBundleExportJobArn: jsii.String("assetBundleExportJobArn"),
//   	AssetBundleExportJobId: jsii.String("assetBundleExportJobId"),
//   }
//
type AssetBundleExportJobReference struct {
	// The ARN of the AssetBundleExportJob resource.
	AssetBundleExportJobArn *string `field:"required" json:"assetBundleExportJobArn" yaml:"assetBundleExportJobArn"`
	// The AssetBundleExportJobId of the AssetBundleExportJob resource.
	AssetBundleExportJobId *string `field:"required" json:"assetBundleExportJobId" yaml:"assetBundleExportJobId"`
}

