package awspanorama


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageLocationProperty := &StorageLocationProperty{
//   	BinaryPrefixLocation: jsii.String("binaryPrefixLocation"),
//   	Bucket: jsii.String("bucket"),
//   	GeneratedPrefixLocation: jsii.String("generatedPrefixLocation"),
//   	ManifestPrefixLocation: jsii.String("manifestPrefixLocation"),
//   	RepoPrefixLocation: jsii.String("repoPrefixLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-panorama-package-storagelocation.html
//
type CfnPackage_StorageLocationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-panorama-package-storagelocation.html#cfn-panorama-package-storagelocation-binaryprefixlocation
	//
	BinaryPrefixLocation *string `field:"optional" json:"binaryPrefixLocation" yaml:"binaryPrefixLocation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-panorama-package-storagelocation.html#cfn-panorama-package-storagelocation-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-panorama-package-storagelocation.html#cfn-panorama-package-storagelocation-generatedprefixlocation
	//
	GeneratedPrefixLocation *string `field:"optional" json:"generatedPrefixLocation" yaml:"generatedPrefixLocation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-panorama-package-storagelocation.html#cfn-panorama-package-storagelocation-manifestprefixlocation
	//
	ManifestPrefixLocation *string `field:"optional" json:"manifestPrefixLocation" yaml:"manifestPrefixLocation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-panorama-package-storagelocation.html#cfn-panorama-package-storagelocation-repoprefixlocation
	//
	RepoPrefixLocation *string `field:"optional" json:"repoPrefixLocation" yaml:"repoPrefixLocation"`
}

