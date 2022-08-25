package awsapigateway


// S3 location of the API definition file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiDefinitionS3Location := &apiDefinitionS3Location{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	version: jsii.String("version"),
//   }
//
type ApiDefinitionS3Location struct {
	// The S3 bucket.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// An optional version.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

