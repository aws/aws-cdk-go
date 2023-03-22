package awspanorama


// A storage location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageLocationProperty := &storageLocationProperty{
//   	binaryPrefixLocation: jsii.String("binaryPrefixLocation"),
//   	bucket: jsii.String("bucket"),
//   	generatedPrefixLocation: jsii.String("generatedPrefixLocation"),
//   	manifestPrefixLocation: jsii.String("manifestPrefixLocation"),
//   	repoPrefixLocation: jsii.String("repoPrefixLocation"),
//   }
//
type CfnPackage_StorageLocationProperty struct {
	// The location's binary prefix.
	BinaryPrefixLocation *string `field:"optional" json:"binaryPrefixLocation" yaml:"binaryPrefixLocation"`
	// The location's bucket.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The location's generated prefix.
	GeneratedPrefixLocation *string `field:"optional" json:"generatedPrefixLocation" yaml:"generatedPrefixLocation"`
	// The location's manifest prefix.
	ManifestPrefixLocation *string `field:"optional" json:"manifestPrefixLocation" yaml:"manifestPrefixLocation"`
	// The location's repo prefix.
	RepoPrefixLocation *string `field:"optional" json:"repoPrefixLocation" yaml:"repoPrefixLocation"`
}

