package awsiot


// The sbom zip archive location of the package version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sbomProperty := &SbomProperty{
//   	S3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-softwarepackageversion-sbom.html
//
type CfnSoftwarePackageVersion_SbomProperty struct {
	// The Amazon S3 location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-softwarepackageversion-sbom.html#cfn-iot-softwarepackageversion-sbom-s3location
	//
	S3Location interface{} `field:"required" json:"s3Location" yaml:"s3Location"`
}

