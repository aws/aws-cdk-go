package previewawsiotmixins


// The artifact location of the package version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   packageVersionArtifactProperty := &PackageVersionArtifactProperty{
//   	S3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-softwarepackageversion-packageversionartifact.html
//
type CfnSoftwarePackageVersionPropsMixin_PackageVersionArtifactProperty struct {
	// The Amazon S3 location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-softwarepackageversion-packageversionartifact.html#cfn-iot-softwarepackageversion-packageversionartifact-s3location
	//
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

