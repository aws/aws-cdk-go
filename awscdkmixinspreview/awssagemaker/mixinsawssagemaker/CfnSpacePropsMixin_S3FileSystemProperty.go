package mixinsawssagemaker


// A custom file system in Amazon S3.
//
// This is only supported in Amazon SageMaker Unified Studio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3FileSystemProperty := &S3FileSystemProperty{
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-s3filesystem.html
//
type CfnSpacePropsMixin_S3FileSystemProperty struct {
	// The Amazon S3 URI that specifies the location in S3 where files are stored, which is mounted within the Studio environment.
	//
	// For example: `s3://<bucket-name>/<prefix>/` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-s3filesystem.html#cfn-sagemaker-space-s3filesystem-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

