package awscdk


// Represents the source for a file asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   fileAssetSource := &FileAssetSource{
//   	SourceHash: jsii.String("sourceHash"),
//
//   	// the properties below are optional
//   	DeployTime: jsii.Boolean(false),
//   	Executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	FileName: jsii.String("fileName"),
//   	Packaging: cdk.FileAssetPackaging_ZIP_DIRECTORY,
//   }
//
type FileAssetSource struct {
	// A hash on the content source.
	//
	// This hash is used to uniquely identify this
	// asset throughout the system. If this value doesn't change, the asset will
	// not be rebuilt or republished.
	SourceHash *string `field:"required" json:"sourceHash" yaml:"sourceHash"`
	// Whether or not the asset needs to exist beyond deployment time;
	//
	// i.e.
	// are copied over to a different location and not needed afterwards.
	// Setting this property to true has an impact on the lifecycle of the asset,
	// because we will assume that it is safe to delete after the CloudFormation
	// deployment succeeds.
	//
	// For example, Lambda Function assets are copied over to Lambda during
	// deployment. Therefore, it is not necessary to store the asset in S3, so
	// we consider those deployTime assets.
	// Default: false.
	//
	DeployTime *bool `field:"optional" json:"deployTime" yaml:"deployTime"`
	// An external command that will produce the packaged asset.
	//
	// The command should produce the location of a ZIP file on `stdout`.
	// Default: - Exactly one of `fileName` and `executable` is required.
	//
	Executable *[]*string `field:"optional" json:"executable" yaml:"executable"`
	// The path, relative to the root of the cloud assembly, in which this asset source resides.
	//
	// This can be a path to a file or a directory, depending on the
	// packaging type.
	// Default: - Exactly one of `fileName` and `executable` is required.
	//
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// Which type of packaging to perform.
	// Default: - Required if `fileName` is specified.
	//
	Packaging FileAssetPackaging `field:"optional" json:"packaging" yaml:"packaging"`
}

