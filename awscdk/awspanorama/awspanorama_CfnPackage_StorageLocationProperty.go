package awspanorama


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
	// `CfnPackage.StorageLocationProperty.BinaryPrefixLocation`.
	BinaryPrefixLocation *string `field:"optional" json:"binaryPrefixLocation" yaml:"binaryPrefixLocation"`
	// `CfnPackage.StorageLocationProperty.Bucket`.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// `CfnPackage.StorageLocationProperty.GeneratedPrefixLocation`.
	GeneratedPrefixLocation *string `field:"optional" json:"generatedPrefixLocation" yaml:"generatedPrefixLocation"`
	// `CfnPackage.StorageLocationProperty.ManifestPrefixLocation`.
	ManifestPrefixLocation *string `field:"optional" json:"manifestPrefixLocation" yaml:"manifestPrefixLocation"`
	// `CfnPackage.StorageLocationProperty.RepoPrefixLocation`.
	RepoPrefixLocation *string `field:"optional" json:"repoPrefixLocation" yaml:"repoPrefixLocation"`
}

