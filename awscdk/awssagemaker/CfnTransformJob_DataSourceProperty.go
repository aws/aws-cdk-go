package awssagemaker


// Describes the location of the channel data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceProperty := &DataSourceProperty{
//   	S3DataSource: &S3DataSourceProperty{
//   		S3DataType: jsii.String("s3DataType"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-datasource.html
//
type CfnTransformJob_DataSourceProperty struct {
	// The S3 location of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-datasource.html#cfn-sagemaker-transformjob-datasource-s3datasource
	//
	S3DataSource interface{} `field:"required" json:"s3DataSource" yaml:"s3DataSource"`
}

