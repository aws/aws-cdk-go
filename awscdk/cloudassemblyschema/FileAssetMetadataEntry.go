package cloudassemblyschema


// Metadata Entry spec for files.
//
// Example:
//   entry := map[string]*string{
//   	"packaging": jsii.String("file"),
//   	"s3BucketParameter": jsii.String("bucket-parameter"),
//   	"s3KeyParamenter": jsii.String("key-parameter"),
//   	"artifactHashParameter": jsii.String("hash-parameter"),
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

