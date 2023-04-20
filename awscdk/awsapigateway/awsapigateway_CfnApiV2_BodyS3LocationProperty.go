package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
// Deprecated: moved to package aws-apigatewayv2.
type CfnApiV2_BodyS3LocationProperty struct {
	// `CfnApiV2.BodyS3LocationProperty.Bucket`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-bucket
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// `CfnApiV2.BodyS3LocationProperty.Etag`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-etag
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// `CfnApiV2.BodyS3LocationProperty.Key`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-key
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// `CfnApiV2.BodyS3LocationProperty.Version`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-version
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

