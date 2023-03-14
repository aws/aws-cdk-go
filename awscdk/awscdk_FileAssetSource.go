// An experiment to bundle the entire CDK into a single module
package awscdk


// Represents the source for a file asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   fileAssetSource := &fileAssetSource{
//   	sourceHash: jsii.String("sourceHash"),
//
//   	// the properties below are optional
//   	executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	fileName: jsii.String("fileName"),
//   	packaging: monocdk.fileAssetPackaging_ZIP_DIRECTORY,
//   }
//
// Experimental.
type FileAssetSource struct {
	// A hash on the content source.
	//
	// This hash is used to uniquely identify this
	// asset throughout the system. If this value doesn't change, the asset will
	// not be rebuilt or republished.
	// Experimental.
	SourceHash *string `field:"required" json:"sourceHash" yaml:"sourceHash"`
	// An external command that will produce the packaged asset.
	//
	// The command should produce the location of a ZIP file on `stdout`.
	// Experimental.
	Executable *[]*string `field:"optional" json:"executable" yaml:"executable"`
	// The path, relative to the root of the cloud assembly, in which this asset source resides.
	//
	// This can be a path to a file or a directory, depending on the
	// packaging type.
	// Experimental.
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// Which type of packaging to perform.
	// Experimental.
	Packaging FileAssetPackaging `field:"optional" json:"packaging" yaml:"packaging"`
}

