package awssagemaker


// The S3 location of the data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DataSourceProperty := &S3DataSourceProperty{
//   	S3DataType: jsii.String("s3DataType"),
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-s3datasource.html
//
type CfnTransformJob_S3DataSourceProperty struct {
	// The data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-s3datasource.html#cfn-sagemaker-transformjob-s3datasource-s3datatype
	//
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// The S3 URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-s3datasource.html#cfn-sagemaker-transformjob-s3datasource-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

