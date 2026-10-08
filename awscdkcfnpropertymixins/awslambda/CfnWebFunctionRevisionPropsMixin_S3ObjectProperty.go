package awslambda


// The S3 location of the function code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3ObjectProperty := &S3ObjectProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   	VersionId: jsii.String("versionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-s3object.html
//
type CfnWebFunctionRevisionPropsMixin_S3ObjectProperty struct {
	// The S3 bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-s3object.html#cfn-lambda-webfunctionrevision-s3object-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 object key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-s3object.html#cfn-lambda-webfunctionrevision-s3object-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The S3 object version ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-s3object.html#cfn-lambda-webfunctionrevision-s3object-versionid
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

