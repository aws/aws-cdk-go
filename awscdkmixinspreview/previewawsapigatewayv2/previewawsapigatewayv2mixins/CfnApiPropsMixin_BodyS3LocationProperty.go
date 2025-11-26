package previewawsapigatewayv2mixins


// The `BodyS3Location` property specifies an S3 location from which to import an OpenAPI definition.
//
// Supported only for HTTP APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bodyS3LocationProperty := &BodyS3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Etag: jsii.String("etag"),
//   	Key: jsii.String("key"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html
//
type CfnApiPropsMixin_BodyS3LocationProperty struct {
	// The S3 bucket that contains the OpenAPI definition to import.
	//
	// Required if you specify a `BodyS3Location` for an API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The Etag of the S3 object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-etag
	//
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// The key of the S3 object.
	//
	// Required if you specify a `BodyS3Location` for an API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The version of the S3 object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

