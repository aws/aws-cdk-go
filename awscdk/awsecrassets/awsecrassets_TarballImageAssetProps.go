package awsecrassets


// Options for TarballImageAsset.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewTarballImageAsset(this, jsii.String("MyBuildImage"), &tarballImageAssetProps{
//   	tarballFile: jsii.String("local-image.tar"),
//   })
//
type TarballImageAssetProps struct {
	// Absolute path to the tarball.
	//
	// It is recommended to to use the script running directory (e.g. `__dirname`
	// in Node.js projects or dirname of `__file__` in Python) if your tarball
	// is located as a resource inside your project.
	TarballFile *string `field:"required" json:"tarballFile" yaml:"tarballFile"`
}

