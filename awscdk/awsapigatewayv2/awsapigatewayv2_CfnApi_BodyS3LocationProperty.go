package awsapigatewayv2


// The `BodyS3Location` property specifies an S3 location from which to import an OpenAPI definition.
//
// Supported only for HTTP APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bodyS3LocationProperty := &bodyS3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	etag: jsii.String("etag"),
//   	key: jsii.String("key"),
//   	version: jsii.String("version"),
//   }
//
type CfnApi_BodyS3LocationProperty struct {
	// The S3 bucket that contains the OpenAPI definition to import.
	//
	// Required if you specify a `BodyS3Location` for an API.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The Etag of the S3 object.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// The key of the S3 object.
	//
	// Required if you specify a `BodyS3Location` for an API.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The version of the S3 object.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

