package cloudassemblyschema


// Metadata Entry spec for files.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileAssetMetadataEntry := &fileAssetMetadataEntry{
//   	artifactHashParameter: jsii.String("artifactHashParameter"),
//   	id: jsii.String("id"),
//   	packaging: jsii.String("packaging"),
//   	path: jsii.String("path"),
//   	s3BucketParameter: jsii.String("s3BucketParameter"),
//   	s3KeyParameter: jsii.String("s3KeyParameter"),
//   	sourceHash: jsii.String("sourceHash"),
//   }
//
type FileAssetMetadataEntry struct {
	// The name of the parameter where the hash of the bundled asset should be passed in.
	ArtifactHashParameter *string `field:"required" json:"artifactHashParameter" yaml:"artifactHashParameter"`
	// Logical identifier for the asset.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Requested packaging style.
	Packaging *string `field:"required" json:"packaging" yaml:"packaging"`
	// Path on disk to the asset.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Name of parameter where S3 bucket should be passed in.
	S3BucketParameter *string `field:"required" json:"s3BucketParameter" yaml:"s3BucketParameter"`
	// Name of parameter where S3 key should be passed in.
	S3KeyParameter *string `field:"required" json:"s3KeyParameter" yaml:"s3KeyParameter"`
	// The hash of the asset source.
	SourceHash *string `field:"required" json:"sourceHash" yaml:"sourceHash"`
}

