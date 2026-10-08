package awslambda


// The code configuration for the revision.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfigProperty := &CodeConfigProperty{
//   	S3Object: &S3ObjectProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//
//   		// the properties below are optional
//   		VersionId: jsii.String("versionId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-codeconfig.html
//
type CfnWebFunctionRevision_CodeConfigProperty struct {
	// The S3 location of the function code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-codeconfig.html#cfn-lambda-webfunctionrevision-codeconfig-s3object
	//
	S3Object interface{} `field:"required" json:"s3Object" yaml:"s3Object"`
}

