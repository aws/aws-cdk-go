package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DataSourceProperty := &S3DataSourceProperty{
//   	CompressionType: jsii.String("compressionType"),
//   	S3DataType: jsii.String("s3DataType"),
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-s3datasource.html
//
type CfnModel_S3DataSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-s3datasource.html#cfn-sagemaker-model-s3datasource-compressiontype
	//
	CompressionType *string `field:"required" json:"compressionType" yaml:"compressionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-s3datasource.html#cfn-sagemaker-model-s3datasource-s3datatype
	//
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-s3datasource.html#cfn-sagemaker-model-s3datasource-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

