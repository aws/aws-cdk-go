package awscdk


// Options for the addDockerImageAsset operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   addDockerImageAssetOptions := &AddDockerImageAssetOptions{
//   	DisplayName: jsii.String("displayName"),
//   }
//
type AddDockerImageAssetOptions struct {
	// A display name to associate with the asset.
	// Default: - No display name.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

