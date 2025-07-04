package awsapigateway


// `S3Location` is a property of the [AWS::ApiGateway::RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource that specifies the Amazon S3 location of a OpenAPI (formerly Swagger) file that defines a set of RESTful APIs in JSON or YAML.
//
// > On January 1, 2016, the Swagger Specification was donated to the [OpenAPI initiative](https://docs.aws.amazon.com/https://www.openapis.org/) , becoming the foundation of the OpenAPI Specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	ETag: jsii.String("eTag"),
//   	Key: jsii.String("key"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-s3location.html
//
type CfnRestApi_S3LocationProperty struct {
	// The name of the S3 bucket where the OpenAPI file is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-s3location.html#cfn-apigateway-restapi-s3location-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The Amazon S3 ETag (a file checksum) of the OpenAPI file.
	//
	// If you don't specify a value, API Gateway skips ETag validation of your OpenAPI file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-s3location.html#cfn-apigateway-restapi-s3location-etag
	//
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// The file name of the OpenAPI file (Amazon S3 object name).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-s3location.html#cfn-apigateway-restapi-s3location-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// For versioning-enabled buckets, a specific version of the OpenAPI file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-s3location.html#cfn-apigateway-restapi-s3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

