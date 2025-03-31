package awsecrassets


// Options for TarballImageAsset.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewTarballImageAsset(this, jsii.String("MyBuildImage"), &TarballImageAssetProps{
//   	TarballFile: jsii.String("local-image.tar"),
//   })
//
type TarballImageAssetProps struct {
	// Absolute path to the tarball.
	//
	// It is recommended to to use the script running directory (e.g. `__dirname`
	// in Node.js projects or dirname of `__file__` in Python) if your tarball
	// is located as a resource inside your project.
	TarballFile *string `field:"required" json:"tarballFile" yaml:"tarballFile"`
	// A display name for this asset.
	//
	// If supplied, the display name will be used in locations where the asset
	// identifier is printed, like in the CLI progress information. If the same
	// asset is added multiple times, the display name of the first occurrence is
	// used.
	//
	// The default is the construct path of the `TarballImageAsset` construct,
	// with respect to the enclosing stack. If the asset is produced by a
	// construct helper function (such as `lambda.Code.fromAssetImage()`), this
	// will look like `MyFunction/AssetImage`.
	//
	// We use the stack-relative construct path so that in the common case where
	// you have multiple stacks with the same asset, we won't show something like
	// `/MyBetaStack/MyFunction/Code` when you are actually deploying to
	// production.
	// Default: - Stack-relative construct path.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

